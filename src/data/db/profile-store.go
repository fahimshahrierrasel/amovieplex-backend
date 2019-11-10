package db

import (
	"amovieplex-backend/src/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

const (
	profileCollectionName = "profiles"
)

type Profile = models.Profile

func CreateProfile(ctx *gin.Context, newProfile Profile) (primitive.ObjectID, error) {
	profileCollection := GetCollection(ctx, profileCollectionName)
	insertResult, err := profileCollection.InsertOne(ctx, newProfile)

	if err != nil {
		log.Printf("Error Inserting New Profile: %v", err)
		return primitive.NilObjectID, err
	}
	log.Printf("Create Profile: Inserted Profile %v", insertResult.InsertedID)

	return insertResult.InsertedID.(primitive.ObjectID), nil
}

func GetProfile(ctx *gin.Context, userID string) (Profile, error) {
	profileCollection := GetCollection(ctx, profileCollectionName)
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Printf("primitive.ObjectIDFromHex ERROR: %v", err)
		return Profile{}, err
	}
	filter := bson.M{"user": userObjectID}
	var profile Profile
	err = profileCollection.FindOne(ctx, filter).Decode(&profile)

	if err != nil {
		log.Printf("Error finding profile from DB: %v", err)
		return Profile{}, err
	}

	return profile, nil
}