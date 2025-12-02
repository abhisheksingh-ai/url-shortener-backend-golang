package controller

import (
	"net/http"
	"urlShortener/internals/dto"
	"urlShortener/internals/service"

	"github.com/gin-gonic/gin"
)

type UrlController struct {
	service service.UrlService
}

func GetUrlController(s service.UrlService) *UrlController {
	return &UrlController{
		service: s,
	}
}

func (c *UrlController) CreateNewShortUrl(ctx *gin.Context) {
	var request dto.UrlDto

	userId := ctx.GetString("userId")
	request.UserId = userId

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload: " + err.Error(),
		})
		return
	}

	response, err := c.service.CreateNewShortUrl(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create short URL: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (c *UrlController) RedirectUrl(ctx *gin.Context) {
	// getting short code
	shortCode := ctx.Param("shortCode")

	requestDto := &dto.UrlDto{
		ShortCode: shortCode,
	}

	resp, err := c.service.RedirectUrl(ctx, requestDto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Redirect(http.StatusFound, resp.OriginalUrl)
}
