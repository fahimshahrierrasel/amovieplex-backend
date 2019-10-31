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
	ratingCollectionName = "ratings"
)

// Rating is the type def of models Ratings
type Rating = models.Rating

// CreateRating will add new rating to mongodb
func CreateRating(ctx *gin.Context, newRating Rating) bool {
	ratingCollection := GetCollection(ctx, ratingCollectionName)
	insertResult, err := ratingCollection.InsertOne(ctx, newRating)
	log.Printf("Create Rating: Inserted Item %v", insertResult.InsertedID)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// GetAllRating will return all rating
func GetAllRating(ctx *gin.Context) []Rating {
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

func GetRating(ctx *gin.Context, ratingID string) (Rating, error) {
	ratingCollection := GetCollection(ctx, ratingCollectionName)
	idPrimitive, err := primitive.ObjectIDFromHex(ratingID)
	if err != nil {
		log.Printf("primitive.ObjectIDFromHex ERROR: %v", err)
	}

	filter := bson.M{"_id": idPrimitive}
	var rating Rating
	err = ratingCollection.FindOne(ctx, filter).Decode(&rating)
	if err != nil {
		log.Printf("Finding Rating from DB Error: %v", err)
	}

	log.Printf("Get Rating: %v", rating)

	return rating, err
}

func UpdateRating(ctx *gin.Context, ratingID string, rating Rating) bool {
	ratingCollection := GetCollection(ctx, ratingCollectionName)
	idPrimitive, err := primitive.ObjectIDFromHex(ratingID)
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	}

	filter := bson.M{"_id": idPrimitive}
	update := bson.D{{Key: "$set", Value: rating}}
	updateResult, err := ratingCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal("ratingCollection.UpdateOne ERROR:", err)
		return false
	}

	log.Printf("Update Rating: %v item(s) Matched and %v item(s) Updated",
		updateResult.MatchedCount, updateResult.ModifiedCount)

	return updateResult.ModifiedCount > 0
}

// SoftDeleteRating will only add deleted_at time which declare its deleted softly
func SoftDeleteRating(ctx *gin.Context, ratingID string) bool {
	ratingCollection := GetCollection(ctx, ratingCollectionName)
	idPrimitive, err := primitive.ObjectIDFromHex(ratingID)
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	}

	filter := bson.M{"_id": idPrimitive}
	update := bson.D{{Key: "$set", Value: bson.M{"deleted_at": time.Now()}}}
	updateResult, err := ratingCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal("ratingCollection.UpdateOne ERROR:", err)
		return false
	}

	log.Printf("SoftDelete Rating: %v item(s) Matched and %v item(s) Soft Deleted",
		updateResult.MatchedCount, updateResult.ModifiedCount)

	return updateResult.ModifiedCount > 0
}

// DeleteRating will delete the given id
func DeleteRating(ctx *gin.Context, ratingID string) bool {
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
	return deleteResult.DeletedCount > 0
}
