package main

import (
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

	router.GET("/users", handler.GetUsers)
	router.GET("/users/:id", handler.GetUserByID)
	router.POST("/users", handler.CreateUser)
	router.DELETE("/users/:id", handler.DeleteUser)

	router.Run(":8080")
}
