package db

import (
	"amovieplex-backend/src/models"
	"github.com/gin-gonic/gin"
	"log"
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
