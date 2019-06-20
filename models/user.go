package models

import (
	"go-mongo/db"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username string        `json:"username" bson:"username"`
	Password string        `json:"password" bson:"password"`
	Fullname string        `json:"fullname" bson:"fullname"`
}

type Update struct {
	Fullname string `json:"fullname" bson:"fullname"`
}

var server = "mongodb://cuong:1@localhost:27017/?authSource=admin"
var dbConnect = db.NewConnection(server)
var collection = dbConnect.Use("go_mongo", "user")

func CheckId(id string) bool {
	if !bson.IsObjectIdHex(id) {
		return false
	}
	return true
}

func FindByUsername(username string) (User, error) {
	object := User{}
	err := collection.Find(bson.M{"username": username}).One(&object)
	return object, err
}

func FindById(id string) (User, error) {
	oid := bson.ObjectIdHex(id)
	user := User{}
	err := collection.FindId(oid).One(&user)

	return user, err
}

func Create(newUser User) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	newUser.Password = string(hash)
	err := collection.Insert(&newUser)
	return err
}

func GetAll() ([]User, error) {
	var listUser []User
	err := collection.Find(nil).All(&listUser)
	return listUser, err
}

func Auth(currentUser User) error {
	user, _ := FindByUsername(currentUser.Username)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(currentUser.Password))
	return err
}

func RemoveById(id string) error {
	oid := bson.ObjectIdHex(id)
	err := collection.RemoveId(oid)
	return err
}

func UpdateById(id string, data Update) error {
	change := bson.M{"$set": bson.M{"fullname": data.Fullname}}
	oid := bson.ObjectIdHex(id)
	err := collection.UpdateId(oid, change)
	return err
}
