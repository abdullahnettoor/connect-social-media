package routes

import (
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/server/handlers"
	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/server/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(
	engine *gin.Engine,
	userHandler *handlers.UserHandler,
	postHandler *handlers.PostHandler,
	commentHandler *handlers.CommentHandler,
) {
	engine.POST("/signUp", userHandler.SignUp)
	engine.POST("/login", userHandler.Login)

	user := engine.Group("/").Use(middlewares.AuthenticateUser)
	user.PATCH("/verifyOtp", userHandler.VerifyOtp)

	user.POST("/newPost", postHandler.CreatePost)

	user.PATCH("/posts/likePost", postHandler.LikePost)
	user.DELETE("/posts/unlikePost", postHandler.UnlikePost)

	user.POST("/commentPost", commentHandler.CreateComment)
	user.DELETE("/deleteComment", commentHandler.DeleteComment)

	user.PATCH("/follow", userHandler.FollowUser)
	user.PATCH("/unfollow", userHandler.UnfollowUser)

	user.GET("/followers", userHandler.GetFollowers)
	user.GET("/following", userHandler.GetFollowing)

	user.GET("/posts", postHandler.GetAllPosts)
	user.GET("/post/:postId/comments", commentHandler.GetCommentsByPostId)
}
