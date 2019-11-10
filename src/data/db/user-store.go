package db

import (
	"amovieplex-backend/src/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

const (
	userCollectionName = "users"
)

type User = models.User

func CreateUser(ctx *gin.Context, newUser User) (primitive.ObjectID, error) {
	userCollection := GetCollection(ctx, userCollectionName)
	insertResult, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		log.Printf("Error Inserting New User: %v", err)
		return primitive.NilObjectID, err
	}
	log.Printf("Create User: Inserted Item %v", insertResult.InsertedID)

	//userID, _ := primitive.ObjectIDFromHex(insertResult.InsertedID.(string))
	return insertResult.InsertedID.(primitive.ObjectID), nil
}

func GetUser(ctx *gin.Context, email string, phoneNo string) (User, error) {
	userCollection := GetCollection(ctx, userCollectionName)
	var filter primitive.M
	if len(email) > 0 {
		filter = bson.M{"email": &email}
	} else {
		filter = bson.M{"phone_number": &phoneNo}
	}
	var user User
	err := userCollection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		log.Printf("Finding User from DB Error: %v", err)
		return User{}, err
	}

	return user, nil
}
