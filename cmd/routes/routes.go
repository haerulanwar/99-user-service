package routes

import (
	"99-user-service/internal/app/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *gin.Engine, userHandler *handlers.UserHandler) {

	app.GET("/users", userHandler.GetAllUsers)
	app.POST("/users", userHandler.CreateUser)
	app.GET("/users/:id", userHandler.GetUserByID)

}
