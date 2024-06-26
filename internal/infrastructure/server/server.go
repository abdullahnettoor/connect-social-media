package server

import (
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/server/handlers"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/server/routes"
	"github.com/gin-gonic/gin"
)

type ServeHttp struct {
	server *gin.Engine
}

func NewServeHttp(
	userHandler *handlers.UserHandler,
	adminHandler *handlers.AdminHandler,
	postHandler *handlers.PostHandler,
	commentHandler *handlers.CommentHandler,
	chatHandler *handlers.ChatHandler,
) *ServeHttp {
	server := gin.New()
	server.Use(gin.Logger())

	server.GET("/check", func(ctx *gin.Context) {
		ctx.String(200, "Server Health OK")
	})

	routes.SetupUserRoutes(server, userHandler, postHandler, commentHandler, chatHandler)
	routes.SetupAdminRoutes(server, adminHandler)

	return &ServeHttp{server}
}

func (s *ServeHttp) Start() {
	s.server.Run("127.0.0.1:9000")
}
