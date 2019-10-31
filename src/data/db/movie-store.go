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

func GetMovie(ctx *gin.Context, movieID string) (Movie, error) {
	movieCollection := GetCollection(ctx, movieCollectionName)
	idPrimitive, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		log.Printf("primitive.ObjectIDFromHex ERROR: %v", err)
		return Movie{}, err
	}

	filter := bson.M{"_id": idPrimitive}
	var movie Movie
	err = movieCollection.FindOne(ctx, filter).Decode(&movie)
	if err != nil {
		log.Printf("Error Finding Movie from DB: %v", err)
		return Movie{}, err
	}
	log.Printf("Get Movie: %v", movie)
	return movie, err
}

// UpdateMovie will update movie for the given id and the update
func UpdateMovie(ctx *gin.Context, movieID string, movieUpdates Movie) bool {
	movieCollection := GetCollection(ctx, movieCollectionName)
	idPrimitive, err := primitive.ObjectIDFromHex(movieID)

	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	}

	filter := bson.M{"_id": idPrimitive}
	update := bson.D{{Key: "$set", Value: movieUpdates}}
	updateResult, err := movieCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal("movieCollection.UpdateOne ERROR:", err)
		return false
	}

	log.Printf("Update Movie: %v item(s) Matched and %v item(s) Updated",
		updateResult.MatchedCount, updateResult.ModifiedCount)

	return updateResult.ModifiedCount > 0
}

// SoftDeleteMovie will only add deleted_at time which declare its deleted softly
func SoftDeleteMovie(ctx *gin.Context, movieID string) bool {
	ratingCollection := GetCollection(ctx, movieCollectionName)
	idPrimitive, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	}

	filter := bson.M{"_id": idPrimitive}
	update := bson.D{{Key: "$set", Value: bson.M{"deleted_at": time.Now()}}}
	updateResult, err := ratingCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal("movieCollection.UpdateOne ERROR:", err)
		return false
	}

	log.Printf("SoftDelete Movie: %v item(s) Matched and %v item(s) Soft Deleted",
		updateResult.MatchedCount, updateResult.ModifiedCount)

	return updateResult.ModifiedCount > 0
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
	return deleteResult.DeletedCount > 0
}
