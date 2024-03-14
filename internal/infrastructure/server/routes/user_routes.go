package routes

import (
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/server/handlers"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(engine *gin.Engine, userHandler *handlers.UserHandler) {
	engine.POST("/signUp", userHandler.SignUp)
	engine.POST("/login", userHandler.Login)
}
