package controller

import (
	"net/http"
	"urlShortener/internals/dto"
	"urlShortener/internals/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func GetNewUserController(s service.UserService) *UserController {
	return &UserController{
		userService: s,
	}
}

func (h *UserController) CreateNewUser(ctx *gin.Context) {
	var request dto.UserDto

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.UserResponseDto{
			Message: "Request body issue",
		})
	}

	//Dto validation
	if request.Password == "" || request.FirstName == "" || request.Email == "" {
		ctx.JSON(http.StatusBadRequest, dto.UserResponseDto{
			Message: "First Name, Email and Password is required field",
		})
		return
	}

	// call the service and send the result
	response, err := h.userService.CreateUser(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.UserResponseDto{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}
