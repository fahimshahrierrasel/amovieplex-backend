package db

import (
	"amovieplex-backend/src/models"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	collection = "ratings"
)

// Rating is the type def of models Ratings
type Rating = models.Rating

// CreateRating will add new rating to mongodb
func CreateRating(context *gin.Context, newRating models.Rating) bool {
	ratingCollection := GetCollection(context, collection)
	insertResult, err := ratingCollection.InsertOne(context, newRating)
	fmt.Println(insertResult)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// GetAllRating will return all rating
func GetAllRating(context *gin.Context) []Rating {
	ratingCollection := GetCollection(context, collection)
	cursor, err := ratingCollection.Find(context, bson.D{})
	var ratings []Rating
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context)
	for cursor.Next(context) {
		result := Rating{}
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		ratings = append(ratings, result)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	log.Printf("All Ratings: %v", ratings)
	return ratings
}

// DeleteRating will delete the given id
func DeleteRating(ctx *gin.Context, ratingID string) bool {
	ratingCollection := GetCollection(ctx, collection)
	idPremitive, err := primitive.ObjectIDFromHex(ratingID)
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	}
	deleteResult, err := ratingCollection.DeleteOne(ctx, bson.M{"_id": idPremitive})
	if err != nil {
		log.Fatal(err)
		return false
	}
	log.Printf("Delete Rating: %v item(s)", deleteResult.DeletedCount)
	if deleteResult.DeletedCount <= 0 {
		return false
	}
	return true
}
