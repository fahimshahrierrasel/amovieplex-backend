package main

import (
	"amovieplex-backend/src/api"
	database "amovieplex-backend/src/data/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variable from .env file
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	app := gin.Default()

	db := database.GetDatabase()
	app.Use(database.Inject(db))

	api.ApplyRoutes(app)

	app.Use(cors.Default())
	// Listen and Server in 0.0.0.0:8080
	app.Run(":8080")
}

