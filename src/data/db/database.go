package db

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetClient return mongo client
func GetDatabase() *mongo.Database {
	// Connection string and db name from the env
	connectionString := os.Getenv("DB_CONN")
	dbName := os.Getenv("DB_NAME")

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database(dbName)
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