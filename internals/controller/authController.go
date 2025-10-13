package controller

import (
	"net/http"
	"urlShortener/internals/service"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService service.AuthService
}

// Object creation
func GetNewAuthController(s service.AuthService) *AuthController {
	return &AuthController{
		authService: s,
	}
}

// Data function
func (h *AuthController) Login(ctx *gin.Context) {
	var loginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalied request boby"})
		return
	}

	token, err := h.authService.Login(ctx, loginReq.Email, loginReq.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid cardnentials"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
