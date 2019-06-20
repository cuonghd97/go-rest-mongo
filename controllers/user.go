package controllers

import (
	"go-mongo/models"
	"log"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

// Create new user
func CreateUser(c *gin.Context) {
	var postData models.User
	err := c.BindJSON(&postData)

	if err != nil {
		c.JSON(200, gin.H{
			"msg": "something went wrong",
		})
		return
	}
	_, err = models.FindByUsername(postData.Username)
	if err == nil {
		c.JSON(200, gin.H{
			"msg": "user already exist",
		})
		return
	}

	models.Create(postData)
	c.JSON(200, gin.H{
		"msg": "create user success",
	})
}

func AllUser(c *gin.Context) {
	listUser, err := models.GetAll()
	if err != nil {
		c.JSON(200, gin.H{
			"msg": "something went wrong",
		})
		return
	}
	c.JSON(200, gin.H{
		"users": listUser,
	})
}

func GetById(c *gin.Context) {
	id := c.Param("id")
	if models.CheckId(id) == false {
		c.JSON(200, gin.H{
			"msg": "wrong id",
		})
		return
	}
	user, err := models.FindById(id)
	if err != nil {
		c.JSON(200, gin.H{
			"msg": "Not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"user": user,
	})
}

func Login(c *gin.Context) {
	var postData models.User
	err := c.BindJSON(&postData)
	if err != nil {
		c.JSON(200, gin.H{
			"msg": "something went wrong",
		})
		return
	}
	err = models.Auth(postData)
	if err != nil {
		c.JSON(200, gin.H{
			"msg": "wrong password or username",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "loggedin",
	})
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	if models.CheckId(id) == false {
		c.JSON(200, gin.H{
			"msg": "wrong id",
		})
		return
	}
	err := models.RemoveById(id)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{
			"msg": "something went wrong",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "Delete success",
	})
}

func Update(c *gin.Context) {
	var postData models.Update
	err := c.BindJSON(&postData)
	if err != nil {
		c.JSON(200, gin.H{
			"msg": "wrong data",
		})
		return
	}
	id := c.Param("id")
	if models.CheckId(id) == false {
		c.JSON(200, gin.H{
			"msg": "wrong id",
		})
		return
	}
	err = models.UpdateById(id, postData)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{
			"msg": "error",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "update",
	})
}
