package db

import (
	"amovieplex-backend/src/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

const (
	genreCollectionName = "genres"
)

type Genre = models.Genre

// CreateGenre will add new genre to mongodb
func CreateGenre(ctx *gin.Context, newGenre Genre) bool {
	genreCollection := GetCollection(ctx, genreCollectionName)
	insertResult, err := genreCollection.InsertOne(ctx, newGenre)
	log.Printf("Create Genre: Inserted Item %v", insertResult.InsertedID)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// GetAllGenre will return all genre
func GetAllGenre(ctx *gin.Context) []Genre {
	genreCollection := GetCollection(ctx, genreCollectionName)
	cursor, err := genreCollection.Find(ctx, bson.D{})
	var genres []Genre
	if err != nil {
		log.Fatal()
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		genre := Genre{}
		err := cursor.Decode(&genre)
		if err != nil {
			log.Fatal(err)
		}
		genres = append(genres, genre)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	log.Printf("All Genres: %v", genres)
	return genres
}

// GetGenre will get genre from the db with the given id
func GetGenre(ctx *gin.Context, genreID string) (Genre, error) {
	genreCollection := GetCollection(ctx, genreCollectionName)
	idPrimitive, err := primitive.ObjectIDFromHex(genreID)
	if err != nil {
		log.Printf("primitive.ObjectIDFromHex ERROR: %v", err)
		return Genre{}, err
	}

	filter := bson.M{"_id": idPrimitive}
	var genre Genre
	err = genreCollection.FindOne(ctx, filter).Decode(&genre)
	if err != nil {
		log.Printf("Error Finding Genre from DB: %v", err)
		return Genre{}, err
	}
	log.Printf("Get Genre: %v", genre)

	return genre, err
}

// UpdateGenre will update the genre with given genreID and genreUpdates
func UpdateGenre(ctx *gin.Context, genreID string, genreUpdates Genre) (bool, error) {
	genreCollection := GetCollection(ctx, genreCollectionName)
	idPrimitive, err := primitive.ObjectIDFromHex(genreID)
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
		return false, err
	}

	filter := bson.M{"_id": idPrimitive}
	update := bson.D{{Key: "$set", Value: genreUpdates}}
	updateResult, err := genreCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal("genreCollection.UpdateOne ERROR:", err)
		return false, err
	}

	log.Printf("Update Genre: %v item(s) Matched and %v item(s) Updated",
		updateResult.MatchedCount, updateResult.ModifiedCount)

	return updateResult.ModifiedCount > 0, nil
}

// SoftDeleteGenre will only add deleted_at time which declare its deleted softly
func SoftDeleteGenre(ctx *gin.Context, genreID string) bool {
	genreCollection := GetCollection(ctx, genreCollectionName)
	idPrimitive, err := primitive.ObjectIDFromHex(genreID)
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	}

	filter := bson.M{"_id": idPrimitive}
	update := bson.D{{Key: "$set", Value: bson.M{"deleted_at": time.Now()}}}
	updateResult, err := genreCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal("genreCollection.UpdateOne ERROR:", err)
		return false
	}

	log.Printf("Soft Delete Genre: %v item(s) Matched and %v item(s) Soft Deleted",
		updateResult.MatchedCount, updateResult.ModifiedCount)

	return updateResult.ModifiedCount > 0
}

// DeleteGenre will delete the given id
func DeleteGenre(ctx *gin.Context, genreID string) bool {
	genreCollection := GetCollection(ctx, genreCollectionName)
	idPrimitive, err := primitive.ObjectIDFromHex(genreID)
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	}
	filter := bson.M{"_id": idPrimitive}
	deleteResult, err := genreCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal("genreCollection.DeleteOne ERROR:", err)
		return false
	}
	log.Printf("Delete Genre: %v item(s)", deleteResult.DeletedCount)

	return deleteResult.DeletedCount > 0
}
