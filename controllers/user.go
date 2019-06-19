package controllers

import (
	"go-mongo/models"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

type UserController struct {
	session *mgo.Session
}

const (
	DB_NAME       = "go_mongo"
	DB_COLLECTION = "user"
)

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// Create new user
func (uc UserController) CreateUser(c *gin.Context) {
	var json models.User
	c.Bind(&json)
	u := uc.create_user(json.Username, json.Password, c)
	if u.Username == json.Username {
		content := gin.H{
			"result": "create success",
		}
		c.Writer.Header().Set("Content-Type", "application/json")
		c.JSON(201, content)
	} else {
		c.JSON(500, gin.H{
			"msg": "Internal server error",
		})
	}
}

func (uc UserController) create_user(Username string, Password string, c *gin.Context) models.User {
	user := models.User{
		Username: Username,
		Password: Password,
	}

	uc.session.DB(DB_NAME).C(DB_COLLECTION).Insert(&user)
	// if err := us.session.DB(DB_NAME).C(DB_COLLECTION).UpdateId(oid, &user); err != nil {

	// }
	return user
}
