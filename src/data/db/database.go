package db

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetClient return mongo client
func GetDatabase() *mongo.Database {
	clientOptions := options.Client().ApplyURI("mongodb://backend-mongodb:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database("test")
	return database
}

func GetCollection(context *gin.Context, collectionName string) *mongo.Collection {
	db := context.MustGet("db").(*mongo.Database)
	return db.Collection(collectionName)
}

func Inject(db *mongo.Database) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("db", db)
		context.Next()
	}
}