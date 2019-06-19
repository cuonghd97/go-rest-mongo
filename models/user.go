package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username string        `json:"username" bson:"username"`
	Password string        `json:"password" bson:"password"`
}
