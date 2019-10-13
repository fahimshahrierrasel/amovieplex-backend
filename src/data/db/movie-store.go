package db

import (
	"log"
	"time"

	"amovieplex-backend/src/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	movieCollectionName = "movies"
)

// Rating is the type def of models Ratings
type Movie = models.Movie

// CreateMovie will add new movie to mongodb
func CreateMovie(ctx *gin.Context, newRating Rating) bool {
	ratingCollection := GetCollection(ctx, ratingCollectionName)
	insertResult, err := ratingCollection.InsertOne(ctx, newRating)
	log.Printf("Create Movie: Inserted Item %v", insertResult.InsertedID)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// GetAllMovie will return all movies
func GetAllMovie(ctx *gin.Context) []Rating {
	ratingCollection := GetCollection(ctx, ratingCollectionName)
	cursor, err := ratingCollection.Find(ctx, bson.D{})
	var ratings []Rating
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
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

// SoftDeleteMovie will only add deleted_at time which declare its deleted softly
func SoftDeleteMovie(ctx *gin.Context, ratingID string) bool {
	ratingCollection := GetCollection(ctx, ratingCollectionName)
	idPrimitive, err := primitive.ObjectIDFromHex(ratingID)
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	}

	filter := bson.M{"_id": idPrimitive}
	update := bson.D{{"$set", bson.M{"deleted_at": time.Now()}}}
	updateResult, err := ratingCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal("ratingCollection.UpdateOne ERROR:", err)
		return false
	}

	log.Printf("SoftDelete Rating: %v item(s) Matched and %v item(s) Soft Deleted",
		updateResult.MatchedCount, updateResult.ModifiedCount)

	if updateResult.ModifiedCount <= 0 {
		return false
	}
	return true
}

// DeleteMovie will delete the given id
func DeleteMovie(ctx *gin.Context, ratingID string) bool {
	ratingCollection := GetCollection(ctx, ratingCollectionName)
	idPrimitive, err := primitive.ObjectIDFromHex(ratingID)
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	}
	filter := bson.M{"_id": idPrimitive}
	deleteResult, err := ratingCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal("genreCollection.DeleteOne ERROR:", err)
		return false
	}
	log.Printf("Delete Rating: %v item(s)", deleteResult.DeletedCount)
	if deleteResult.DeletedCount <= 0 {
		return false
	}
	return true
}
