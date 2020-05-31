package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* User Model */
type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string `bson:"name" json:"name,omitempty"`
	LastName string `bson:"lastJame" json:"lastName,omitempty"`
	BirthDay time.Time `bson:"birthDay" json:"birthDay,omitempty"`
	Email string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password,omitempty"`
	Avatar string `bson:"avatar" json:"avatar,omitempty"`
	Banner string `bson:"banner" json:"banner,omitempty"`
	Bio string `bson:"bio" json:"bio,omitempty"`
	Ubication string `bson:"ubication" json:"ubication,omitempty"`
	WebSite string `bson:"website" json:"webSite,omitempty"`
}