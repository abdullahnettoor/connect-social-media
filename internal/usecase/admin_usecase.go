package usecase

import (
	"context"
	"net/http"
	"time"

	"github.com/abdullahnettoor/connect-social-media/internal/domain/constants"
	e "github.com/abdullahnettoor/connect-social-media/internal/domain/error"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/model/req"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/model/res"
	"github.com/abdullahnettoor/connect-social-media/internal/repo"
	"github.com/abdullahnettoor/connect-social-media/pkg/helper"
	jwttoken "github.com/abdullahnettoor/connect-social-media/pkg/jwt"
	"github.com/spf13/viper"
)

type AdminUseCase struct {
	repo     *repo.AdminRepository
	userRepo *repo.UserRepository
}

func NewAdminUseCase(repo *repo.AdminRepository, userRepo *repo.UserRepository) *AdminUseCase {
	return &AdminUseCase{repo, userRepo}
}

func (uc *AdminUseCase) Login(ctx context.Context, req *req.AdminLoginReq) *res.AdminLoginRes {
	admin, err := uc.repo.FindAdminByEmail(ctx, req.Email)

	switch {
	case err == e.ErrAdminNotFound:
		return &res.AdminLoginRes{CommonRes: res.CommonRes{
			Code:    http.StatusNotFound,
			Message: "Admin not found",
			Error:   err.Error()},
		}
	case err != nil:
		return &res.AdminLoginRes{CommonRes: res.CommonRes{
			Code:    http.StatusInternalServerError,
			Message: "server error",
			Error:   err.Error()},
		}
	}

	token, err := jwttoken.CreateToken(viper.GetString("JWT_SECRET"), "admin", time.Hour*24, admin)
	if err != nil {
		return &res.AdminLoginRes{CommonRes: res.CommonRes{
			Code:    http.StatusInternalServerError,
			Message: "failed to generate token",
			Error:   err.Error()},
		}
	}
	admin.Password = ""

	return &res.AdminLoginRes{CommonRes: res.CommonRes{
		Code:    http.StatusOK,
		Message: "User logged in successfully"},
		Token: token,
		Admin: *admin,
	}
}

func (uc *AdminUseCase) GetUser(ctx context.Context, req *req.UserId) *res.CommonRes {
	user, err := uc.userRepo.FindUserByUserId(ctx, req.UserID)
	if err != nil {
		return &res.CommonRes{
			Code:    http.StatusInternalServerError,
			Message: "failed to generate token",
			Error:   err.Error(),
		}
	}
	return &res.CommonRes{
		Code:   200,
		Result: user,
	}
}

func (uc *AdminUseCase) BlockUser(ctx context.Context, req *req.UserId) *res.CommonRes {
	user, err := uc.userRepo.UpdateUserStatus(
		ctx,
		req.UserID,
		string(constants.UserStatusBlocked),
		helper.CurrentIsoDateTimeString(),
	)
	if err != nil {
		return &res.CommonRes{
			Code:    http.StatusInternalServerError,
			Message: "failed to generate token",
			Error:   err.Error(),
		}
	}
	user.Password = ""
	return &res.CommonRes{
		Code:   200,
		Result: user,
	}
}

func (uc *AdminUseCase) UnblockUser(ctx context.Context, req *req.UserId) *res.CommonRes {
	user, err := uc.userRepo.UpdateUserStatus(
		ctx,
		req.UserID,
		string(constants.UserStatusActive),
		helper.CurrentIsoDateTimeString(),
	)
	if err != nil {
		return &res.CommonRes{
			Code:    http.StatusInternalServerError,
			Message: "failed to generate token",
			Error:   err.Error(),
		}
	}
	user.Password = ""
	return &res.CommonRes{
		Code:   200,
		Result: user,
	}
}
