package db

import (
	"context"
	"time"
	"github.com/AdrianMendez1199/red-social-go/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Register(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()
	
	db := MongoCN.Database("twitter")
	col := db.Collection("users")

	u.Password, _ = EncriptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	objID, _ :=  result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}