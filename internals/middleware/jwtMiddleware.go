package middleware

import (
	"net/http"
	"strings"
	"urlShortener/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			utils.InitLogger().Info("Authorization header or token is missing")
			ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "Missing token/ Authorization header"})
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ValidatToken(tokenString)
		if err != nil {
			utils.InitLogger().Error(err.Error())
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		ctx.Set("userId", claims.UserID)
		ctx.Set("email", claims.Email)
	}
}
