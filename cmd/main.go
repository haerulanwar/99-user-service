package main

import (
	"99-user-service/cmd/routes"
	"99-user-service/config"
	"99-user-service/database"
	"99-user-service/internal/app/handlers"
	"99-user-service/internal/app/repositories"
	"99-user-service/internal/app/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	config.Load()

	database.InitDB()

	app := gin.Default()

	userRepo := repositories.NewUserRepository(database.DB)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	routes.SetupRoutes(app, userHandler)
	port := config.GetEnv("PORT", "8081")

	log.Printf("API Gateway running on port %s", port)
	app.Run(":" + port)
}
