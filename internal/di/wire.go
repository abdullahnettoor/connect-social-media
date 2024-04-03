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
	imgupload "github.com/abdullahnettoor/connect-social-media/pkg/image_uploader"

	"github.com/google/wire"
)

func InitializeAPI(cfg *config.Config) (*server.ServeHttp, error) {
	wire.Build(
		server.NewServeHttp,
		db.ConnectDb,
		imgupload.ConnectCloudinary,
		repo.NewUserRepository,
		usecase.NewUserUseCase,
		handlers.NewUserHandler,
		repo.NewAdminRepository,
		usecase.NewAdminUseCase,
		handlers.NewAdminHandler,
		repo.NewPostRepository,
		repo.NewContentRepository,
		usecase.NewPostUseCase,
		handlers.NewPostHandler,
		repo.NewCommentRepository,
		usecase.NewCommentUseCase,
		handlers.NewCommentHandler,
		repo.NewChatRepository,
		usecase.NewChatUseCase,
		handlers.NewChatHandler,
	)

	return &server.ServeHttp{}, nil
}
