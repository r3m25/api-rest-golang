package repository

import (
	"context"
	"github.com/r3m25/api-rest-golang/bd/configuration"
	"github.com/r3m25/api-rest-golang/bd/encrypt"
	"github.com/r3m25/api-rest-golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func SaveUser(user models.User) (string, bool, error)  {

	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := configuration.Connection().Database("golang-bd")
	collection := db.Collection("user")

	user.Password, _ = encrypt.PasswordEncrypt(user.Password)

	result, err := collection.InsertOne(ctx, user)

	if err!= nil {
		return "", false, err
	}

	objectID, _ := result.InsertedID.(primitive.ObjectID)

	return objectID.String(), true, nil
}

func CheckIfExistUser(mail string) (models.User, bool, string) {

	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := configuration.Connection().Database("golang-bd")
	collection := db.Collection("user")

	condition := bson.M{"mail": mail}

	var result models.User

	err := collection.FindOne(ctx, condition).Decode(&result)

	id := result.ID.Hex()

	if err != nil {
		return result, false, id
	}
	return result, true, id
}

