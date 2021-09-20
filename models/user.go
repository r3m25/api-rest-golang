package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID 			primitive.ObjectID 	`bson:"_id,omitempty" json:"id"`
	Name		string 				`bson:"name" json:"name,omitempty"`
	LastName 	string 				`bson:"lastName" json:"lastName,omitempty"`
	Birthday 	time.Time 			`bson:"birthday" json:"birthday,omitempty"`
	Mail 		string 				`bson:"mail" json:"mail"`
	Password 	string 				`bson:"password" json:"password,omitempty"`
	Avatar		string				`bson:"avatar" json:"avatar,omitempty"`
	Banner		string				`bson:"banner" json:"banner,omitempty"`
	Biography	string				`bson:"biography" json:"biography,omitempty"`
	Location	string				`bson:"location" json:"location,omitempty"`
	WebSite		string				`bson:"webSite" json:"webSite,omitempty"`
}
