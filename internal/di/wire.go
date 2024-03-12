//go:build wireinject
// +build wireinject

package di

import (
	"github.com/abdullahnettoor/connect-social-media/internal/config"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/db"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/server"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/server/handlers"
	"github.com/abdullahnettoor/connect-social-media/internal/repo"
	"github.com/abdullahnettoor/connect-social-media/internal/usecase"

	"github.com/google/wire"
)

func InitializeAPI(cfg *config.Config) (*server.ServeHttp, error) {
	wire.Build(
		server.NewServeHttp,
		db.ConnectDb,
		repo.NewUserRepository,
		usecase.NewUserUsecase,
		handlers.NewUserHandler,
		repo.NewAdminRepository,
		usecase.NewAdminUsecase,
		handlers.NewAdminHandler,
	)

	return &server.ServeHttp{}, nil
}
