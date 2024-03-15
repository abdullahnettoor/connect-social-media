package usecase

import (
	"context"
	"net/http"
	"time"

	e "github.com/abdullahnettoor/connect-social-media/internal/domain/error"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/model/req"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/model/res"
	"github.com/abdullahnettoor/connect-social-media/internal/repo"
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

func (uc *AdminUseCase) Login(ctx context.Context, req *req.AdminLoginReq) (*res.AdminLoginRes, error) {
	admin, err := uc.repo.FindAdminByEmail(ctx, req.Email)

	switch {
	case err == e.ErrAdminNotFound:
		return &res.AdminLoginRes{Code: http.StatusNotFound, Message: "Admin not found"}, err
	case err != nil:
		return &res.AdminLoginRes{Code: http.StatusInternalServerError, Message: "server error", Error: err}, err
	}

	token, err := jwttoken.CreateToken(viper.GetString("JWT_SECRET"), "admin", time.Hour*24, admin)
	if err != nil {
		return &res.AdminLoginRes{Code: http.StatusInternalServerError, Message: "failed to generate token", Error: err}, err
	}

	admin.Password = ""

	return &res.AdminLoginRes{Code: http.StatusOK, Message: "User logged in successfully", Token: token, Admin: *admin}, nil
}

func (uc *AdminUseCase) GetUser(ctx context.Context, req *req.UserId) (*res.CommonRes, error) {
	user, err := uc.userRepo.FindUserByUserId(ctx, req.UserID)
	if err != nil {
		return &res.CommonRes{Code: http.StatusInternalServerError, Message: "failed to generate token", Error: err}, err
	}
	user.Password = ""
	return &res.CommonRes{Code: 200, Result: user}, nil
}

func (uc *AdminUseCase) BlockUser(ctx context.Context, req *req.UserId) (*res.CommonRes, error) {
	user, err := uc.userRepo.UpdateUserStatus(ctx, req.UserID, "BLOCKED", time.Now().Format(time.RFC3339))
	if err != nil {
		return &res.CommonRes{Code: http.StatusInternalServerError, Message: "failed to generate token", Error: err}, err
	}
	user.Password = ""
	return &res.CommonRes{Code: 200, Result: user}, nil
}

func (uc *AdminUseCase) UnblockUser(ctx context.Context, req *req.UserId) (*res.CommonRes, error) {
	user, err := uc.userRepo.UpdateUserStatus(ctx, req.UserID, "ACTIVE", time.Now().Format(time.RFC3339))
	if err != nil {
		return &res.CommonRes{Code: http.StatusInternalServerError, Message: "failed to generate token", Error: err}, err
	}
	user.Password = ""
	return &res.CommonRes{Code: 200, Result: user}, nil
}
