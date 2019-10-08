package db

import (
	"amovieplex-backend/src/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	collection = "rating"
)

// CreateRating will add new rating to mongodb
func CreateRating(context *gin.Context, newRating models.Rating) bool {
	ratingCollection := GetCollection(context, collection)
	insertResult, err := ratingCollection.InsertOne(context, newRating)
	fmt.Println(insertResult)
	if err != nil {
		return false
	}
	return true
}



func GetAllRating(context *gin.Context) {
	ratingCollection := GetCollection(context, collection)
	cursor, _ := ratingCollection.Find(context, bson.D{})
	fmt.Println(cursor)
}
