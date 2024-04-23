package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/abdullahnettoor/connect-social-media/internal/infrastructure/model/res"
	jwttoken "github.com/abdullahnettoor/connect-social-media/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Authorize admin
func AuthenticateAdmin(ctx *gin.Context) {
	log.Println("MW: Authorizing Admin")

	tokenString := strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")

	var secretKey = viper.GetString("JWT_SECRET")
	// Check if it is admin
	isValid, claims := jwttoken.IsValidToken(secretKey, tokenString)
	if !isValid {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	_ = claims.(*jwttoken.CustomClaims).Model
	role := claims.(*jwttoken.CustomClaims).Role
	if role != "admin" {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}

	log.Println("MW: Admin Authorized")
	ctx.Next()
}

// Authorize user
func AuthenticateUser(ctx *gin.Context) {
	log.Println("MW: Authorizing User")

	tokenString := strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")

	var secretKey = viper.GetString("JWT_SECRET")

	isValid, claims := jwttoken.IsValidToken(secretKey, tokenString)
	if !isValid {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user := claims.(*jwttoken.CustomClaims).Model.(map[string]any)
	role := claims.(*jwttoken.CustomClaims).Role
	if role != "user" {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}
	v, ok := user["status"].(string)
	if ok {
		switch v {
		case "PENDING":
			ctx.JSON(http.StatusForbidden, res.CommonRes{
				Code:    http.StatusForbidden,
				Message: "Verify OTP to Login",
			})
			return
		case "BLOCKED":
			ctx.JSON(http.StatusForbidden, res.CommonRes{
				Code:    http.StatusForbidden,
				Message: "Your profile is blocked",
			})
			return
		}
	}

	ctx.Set("user", user)

	log.Println("MW: User Authorized")
	ctx.Next()
}
