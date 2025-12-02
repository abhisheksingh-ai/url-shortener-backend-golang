package main

import (
	"log"
	"urlShortener/internals/controller"
	"urlShortener/internals/middleware"
	"urlShortener/internals/model"
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

	// auto table creation
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.URL{})

	// authmiddleware
	authMiddleware := middleware.AuthMiddleware()

	//url
	repo := repository.GetUrlRepo(logger, db)
	svc := service.GetUrlService(logger, repo)
	ctrl := controller.GetUrlController(svc)

	//user
	userRepo := repository.GetUserRepository(db, logger)
	userService := service.GetNewService(userRepo, logger)
	userController := controller.GetNewUserController(userService)

	// auth
	authService := service.GetAuthService(userRepo, logger)
	authController := controller.GetNewAuthController(authService)

	//Singup and Login
	//Routes for user
	r.POST("/signup", userController.CreateNewUser) // working this will only create the user
	r.POST("/login", authController.Login)

	//Routes for url
	// protected urls
	protected := r.Group("/api", authMiddleware)
	{
		protected.POST("/shorten", ctrl.CreateNewShortUrl)
		protected.GET("/:shortCode", ctrl.RedirectUrl)
	}

	r.Run(":8080")
}
