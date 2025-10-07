package main

import (
	"log"
	"urlShortener/internals/controller"
	"urlShortener/internals/repository"
	"urlShortener/internals/service"
	"urlShortener/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := gin.Default()

	// Initialize logger, db connection, repo, service, controller
	logger := utils.GetLogger()
	db := utils.GetDbConnection()

	repo := repository.GetUrlRepo(logger, db)
	svc := service.GetUrlService(logger, repo)
	ctrl := controller.GetUrlController(svc)

	r.POST("/api/v1/shorten", ctrl.CreateNewShortUrl)

	r.Run(":8080")
}
