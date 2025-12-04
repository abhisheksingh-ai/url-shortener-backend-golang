package controller

import (
	"log/slog"
	"net/http"
	"urlShortener/internals/dto"
	"urlShortener/internals/service"

	"github.com/gin-gonic/gin"
)

type UrlController struct {
	service service.UrlService
	logger  *slog.Logger
}

func GetUrlController(s service.UrlService, l *slog.Logger) *UrlController {
	return &UrlController{
		service: s,
		logger:  l,
	}
}

func (c *UrlController) CreateNewShortUrl(ctx *gin.Context) {
	var request dto.UrlDto

	userId := ctx.GetString("userId")
	request.UserId = userId

	if err := ctx.ShouldBindJSON(&request); err != nil {
		c.logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload: " + err.Error(),
		})
		return
	}

	response, err := c.service.CreateNewShortUrl(ctx, &request)
	if err != nil {
		c.logger.Error(err.Error())
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

	c.logger.Info("", slog.Any(
		"shortCode", shortCode,
	))

	requestDto := &dto.UrlDto{
		ShortCode: shortCode,
	}

	resp, err := c.service.RedirectUrl(ctx, requestDto)
	if err != nil {
		c.logger.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Redirect(http.StatusFound, resp.OriginalUrl)
}
