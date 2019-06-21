package main

import (
	"go-mongo/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/users", controllers.CreateUser)
	router.GET("/users/:id", controllers.GetById)
	router.GET("/list-users", controllers.AllUser)
	router.POST("/login", controllers.Login)
	router.DELETE("/users/:id", controllers.Delete)
	router.PATCH("/users/:id", controllers.Update)
	router.Run()
}
