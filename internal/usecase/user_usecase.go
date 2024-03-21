package usecase

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/abdullahnettoor/connect-social-media/internal/domain/constants"
	"github.com/abdullahnettoor/connect-social-media/internal/domain/entity"
	e "github.com/abdullahnettoor/connect-social-media/internal/domain/error"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/model/req"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/model/res"
	"github.com/abdullahnettoor/connect-social-media/internal/repo"
	"github.com/abdullahnettoor/connect-social-media/pkg/emailer"
	"github.com/abdullahnettoor/connect-social-media/pkg/helper"
	jwttoken "github.com/abdullahnettoor/connect-social-media/pkg/jwt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type UserUseCase struct {
	repo *repo.UserRepository
}

func NewUserUseCase(repo *repo.UserRepository) *UserUseCase {
	return &UserUseCase{repo}
}

func (uc *UserUseCase) SignUp(ctx context.Context, req *req.SignUpReq) *res.SignUpRes {

	hashed, err := helper.HashPassword(req.Password)
	if err != nil {
		msg := "Error Occurred while hashing password"
		log.Println(msg)
		return &res.SignUpRes{
			CommonRes: res.CommonRes{
				Code:    http.StatusInternalServerError,
				Message: msg,
				Error:   err.Error(),
			},
		}
	}

	otp, _ := helper.GenerateOTP()

	user := &entity.User{
		ID:          uuid.NewString(),
		FullName:    req.FullName,
		Username:    req.Username,
		Email:       req.Email,
		Password:    hashed,
		AccountType: req.AccType,
		Status:      constants.UserStatus(otp),
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}

	user, err = uc.repo.CreateUser(ctx, user)
	switch err {
	case e.ErrUsernameConflict, e.ErrEmailAndUsernameConflict, e.ErrEmailConflict:
		return &res.SignUpRes{
			CommonRes: res.CommonRes{
				Code:    http.StatusConflict,
				Message: "User already exist",
				Error:   err.Error(),
			},
		}
	}
	if err != nil {
		return &res.SignUpRes{
			CommonRes: res.CommonRes{
				Code:    http.StatusInternalServerError,
				Message: "server error",
				Error:   err.Error(),
			},
		}
	}

	sender := viper.GetString("SMTP_EMAIL")
	pwd := viper.GetString("SMTP_PASSWORD")
	if err := emailer.SendOtp(
		sender,
		user.Email,
		pwd,
		otp,
		"Connectr - OTP Verification",
	); err != nil {
		uc.repo.RemoveUserByEmail(ctx, req.Email)
		msg := "Error Occurred while sending otp"
		log.Println(msg)
		return &res.SignUpRes{
			CommonRes: res.CommonRes{
				Code:    http.StatusInternalServerError,
				Message: msg,
				Error:   err.Error(),
			},
		}
	}

	token, err := jwttoken.CreateToken(viper.GetString("JWT_SECRET"), "user", time.Hour*24, user)
	if err != nil {
		uc.repo.RemoveUserByEmail(ctx, req.Email)
		return &res.SignUpRes{
			CommonRes: res.CommonRes{
				Code:    http.StatusInternalServerError,
				Message: "failed to generate token",
				Error:   err.Error(),
			},
		}
	}

	return &res.SignUpRes{
		CommonRes: res.CommonRes{
			Code:    http.StatusOK,
			Message: "Verify OTP in your email",
		}, Token: token,
	}
}

func (uc *UserUseCase) VerifyOtp(ctx context.Context, req *req.VerifyOtp) *res.CommonRes {

	user, err := uc.repo.FindUserByUserId(ctx, req.UserID)
	switch {
	case err == e.ErrUserNotFound:
		return &res.CommonRes{
			Code:    http.StatusBadRequest,
			Message: "user not found with provided id",
			Error:   err.Error(),
		}
	case err != nil:
		log.Println(err)
		return &res.CommonRes{
			Code:    http.StatusInternalServerError,
			Message: "server error",
			Error:   err.Error(),
		}
	}

	if user.Status != constants.UserStatus(req.Otp) {
		return &res.CommonRes{
			Code:    http.StatusBadRequest,
			Message: "The OTP you entered is invalid",
			Error:   e.ErrInvalidOtp.Error(),
		}
	}

	otpSentAt, err := time.Parse(time.RFC3339, user.UpdatedAt)
	if err != nil {
		log.Println(err)
		return &res.CommonRes{
			Code:    http.StatusBadRequest,
			Error:   err.Error(),
			Message: "time parsing error",
		}
	}
	if time.Now().After(otpSentAt.Add(time.Minute * 5)) {
		return &res.CommonRes{
			Code:    http.StatusBadRequest,
			Message: "Your OTP has expired",
			Error:   e.ErrOtpTimeOut.Error(),
		}
	}

	user, err = uc.repo.UpdateUserStatus(
		ctx,
		req.UserID,
		string(constants.UserStatusActive),
		helper.CurrentIsoDateTimeString())
	if err != nil {
		log.Println(err)
		return &res.CommonRes{
			Code:    http.StatusInternalServerError,
			Message: "Error Occurred while updating user status",
			Error:   err.Error(),
		}
	}

	token, err := jwttoken.CreateToken(viper.GetString("JWT_SECRET"), "user", time.Hour*24, user)
	if err != nil {
		return &res.CommonRes{
			Code:    http.StatusInternalServerError,
			Message: "failed to generate token",
			Error:   err.Error(),
		}
	}

	return &res.CommonRes{
		Code:    http.StatusOK,
		Message: "OTP Verification Successful",
		Result:  map[string]any{"user": user, "token": token},
	}
}

func (uc *UserUseCase) Login(ctx context.Context, req *req.LoginReq) *res.LoginRes {
	user, err := uc.repo.FindUserByUsername(ctx, req.Username)

	switch {
	case err == e.ErrUserNotFound:
		return &res.LoginRes{CommonRes: res.CommonRes{
			Code:    http.StatusNotFound,
			Message: "User not found",
			Error:   err.Error(),
		},
		}
	case err != nil:
		return &res.LoginRes{CommonRes: res.CommonRes{
			Code:    http.StatusInternalServerError,
			Message: "server error",
			Error:   err.Error(),
		},
		}
	}

	if err := helper.CompareHashedPassword(user.Password, req.Password); err != nil {
		return &res.LoginRes{CommonRes: res.CommonRes{
			Code:    http.StatusUnauthorized,
			Message: "Invalid Password",
			Error:   err.Error(),
		},
		}
	}
	user.Password = ""

	token, err := jwttoken.CreateToken(
		viper.GetString("JWT_SECRET"),
		"user",
		time.Hour*24,
		user)
	if err != nil {
		return &res.LoginRes{CommonRes: res.CommonRes{
			Code:    http.StatusInternalServerError,
			Message: "failed to generate token",
			Error:   err.Error(),
		},
		}
	}

	return &res.LoginRes{CommonRes: res.CommonRes{
		Code:    http.StatusOK,
		Message: "User logged in successfully",
	},
		Token: token,
		User:  *user,
	}
}
