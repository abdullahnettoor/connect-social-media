package usecase

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/abdullahnettoor/connect-social-media/internal/domain/entity"
	e "github.com/abdullahnettoor/connect-social-media/internal/domain/error"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/model/req"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/model/res"
	"github.com/abdullahnettoor/connect-social-media/internal/repo"
	"github.com/abdullahnettoor/connect-social-media/pkg/helper"
	jwttoken "github.com/abdullahnettoor/connect-social-media/pkg/jwt"
	"github.com/spf13/viper"
)

type UserUseCase struct {
	repo *repo.UserRepository
}

func NewUserUseCase(repo *repo.UserRepository) *UserUseCase {
	return &UserUseCase{repo}
}

func (uc *UserUseCase) SignUp(ctx context.Context, req *req.SignUpReq) (*entity.User, error) {
	hashed, err := helper.HashPassword(req.Password)
	if err != nil {
		log.Println("Error Occurred while hashing password")
		return nil, err
	}

	user := &entity.User{FullName: req.FullName, Username: req.Username, Email: req.Email, Password: hashed}
	user, err = uc.repo.CreateUser(ctx, user)
	if err != nil {
		log.Println("Error Occurred while creating user")
		return nil, err
	}
	return user, nil

}

func (uc *UserUseCase) Login(ctx context.Context, req *req.LoginReq) (*res.LoginRes, error) {
	user, err := uc.repo.FindUserByUsername(ctx, req.Username)

	switch {
	case err == e.ErrUserNotFound:
		return &res.LoginRes{Code: http.StatusNotFound, Message: "User not found"}, err
	case err != nil:
		return &res.LoginRes{Code: http.StatusInternalServerError, Message: "server error", Error: err}, err
	}

	if err := helper.CompareHashedPassword(user.Password, req.Password); err != nil {
		return &res.LoginRes{Code: http.StatusUnauthorized, Message: "Invalid Password"}, err
	}

	token, err := jwttoken.CreateToken(viper.GetString("JWT_SECRET"), "user", time.Hour*24, user)
	if err != nil {
		return &res.LoginRes{Code: http.StatusInternalServerError, Message: "failed to generate token", Error: err}, err
	}

	user.Password = ""

	return &res.LoginRes{Code: http.StatusOK, Message: "User logged in successfully", Token: token, User: *user}, nil
}
