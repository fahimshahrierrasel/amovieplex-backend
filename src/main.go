package main

import (
	"amovieplex-backend/src/api"
	database "amovieplex-backend/src/data/db"
	"amovieplex-backend/src/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// Load environment variable from .env file
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading env file %v", err)
	}

	app := gin.Default()

	db := database.GetDatabase()
	app.Use(database.Inject(db))
	app.Use(middlewares.JWTMiddleware())

	api.ApplyRoutes(app)

	app.Use(cors.Default())
	// Listen and Server in 0.0.0.0:8080
	_ = app.Run(":8080")
}

