package routes

import (
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/server/handlers"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/server/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupAdminRoutes(engine *gin.Engine, adminHandler *handlers.AdminHandler) {
	engine.POST("/adminLogin", adminHandler.Login)
	admin := engine.Group("/admin")
	admin.Use(middlewares.AuthenticateAdmin)
	admin.PATCH("/users/block", adminHandler.BlockUser)
	admin.PATCH("/users/unblock", adminHandler.UnblockUser)
}
