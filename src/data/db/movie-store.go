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
func CreateMovie(ctx *gin.Context, newMovie Movie) bool {
	movieCollection := GetCollection(ctx, movieCollectionName)
	insertResult, err := movieCollection.InsertOne(ctx, newMovie)
	log.Printf("Create Movie: Inserted Item %v", insertResult.InsertedID)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// GetAllMovie will return all movies
func GetAllMovie(ctx *gin.Context) []Movie {
	movieCollection := GetCollection(ctx, movieCollectionName)
	cursor, err := movieCollection.Find(ctx, bson.D{})
	var movies []Movie
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		result := Movie{}
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, result)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	log.Printf("All Movies: %v", movies)
	return movies
}

// SoftDeleteMovie will only add deleted_at time which declare its deleted softly
func SoftDeleteMovie(ctx *gin.Context, movieID string) bool {
	ratingCollection := GetCollection(ctx, movieCollectionName)
	idPrimitive, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	}

	filter := bson.M{"_id": idPrimitive}
	update := bson.D{{"$set", bson.M{"deleted_at": time.Now()}}}
	updateResult, err := ratingCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal("movieCollection.UpdateOne ERROR:", err)
		return false
	}

	log.Printf("SoftDelete Movie: %v item(s) Matched and %v item(s) Soft Deleted",
		updateResult.MatchedCount, updateResult.ModifiedCount)

	if updateResult.ModifiedCount <= 0 {
		return false
	}
	return true
}

// DeleteMovie will delete the given id
func DeleteMovie(ctx *gin.Context, movieID string) bool {
	movieCollection := GetCollection(ctx, movieCollectionName)
	idPrimitive, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	}
	filter := bson.M{"_id": idPrimitive}
	deleteResult, err := movieCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal("movieCollection.DeleteOne ERROR:", err)
		return false
	}
	log.Printf("Delete Movie: %v item(s)", deleteResult.DeletedCount)
	if deleteResult.DeletedCount <= 0 {
		return false
	}
	return true
}
