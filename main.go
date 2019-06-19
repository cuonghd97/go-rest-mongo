package main

import (
	"./controllers"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func main() {
	uc := controllers.NewUserController(getSession())

	router := gin.Default()
	router.POST("/users", uc.CreateUser)

	router.Run()
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return s
}
