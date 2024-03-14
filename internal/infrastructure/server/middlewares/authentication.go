package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	jwttoken "github.com/abdullahnettoor/connect-social-media/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Authorize admin
func AuthenticateAdmin(ctx *gin.Context) {
	fmt.Println("MW: Authorizing Admin")

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

	fmt.Println("MW: Admin Authorized")
	ctx.Next()
}
