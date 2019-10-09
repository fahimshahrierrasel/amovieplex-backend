package db

import (
	"amovieplex-backend/src/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

var (
	collection = "ratings"
)

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
func GetAllRating(context *gin.Context) {
	ratingCollection := GetCollection(context, collection)
	fmt.Println(ratingCollection)
	cursor, err := ratingCollection.Find(context, bson.D{})
	if err != nil { log.Fatal(err) }
	defer cursor.Close(context)
	for cursor.Next(context){
		result := Rating{}
		err := cursor.Decode(&result)
		if err != nil {log.Fatal(err)}
		fmt.Println(result)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	return
}
