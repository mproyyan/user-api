package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mproyyan/user-api/internal/handlers"
	"github.com/mproyyan/user-api/internal/repositories"
	"github.com/mproyyan/user-api/internal/services"
)

func main() {
	repo := repositories.NewInMemoryUserRepository()
	service := services.NewUserService(repo)
	handler := handlers.NewUserHandler(service)

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome 👋",
			"endpoints": []string{
				"GET /users        		- Get list of users",
				"POST /users       		- Create new user",
				"GET /users/:id    		- Get user by ID",
				"DELETE /users/:id 		- Delete user by ID",
			},
		})
	})

	router.GET("/users", handler.GetUsers)
	router.GET("/users/:id", handler.GetUserByID)
	router.POST("/users", handler.CreateUser)
	router.DELETE("/users/:id", handler.DeleteUser)

	host := os.Getenv("APP_HOST")
	if host == "" {
		host = "127.0.0.1:8080" // default
	}

	router.Run(host)
}
