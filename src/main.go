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
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	db := database.GetDatabase()
	app.Use(database.Inject(db))
	app.Use(middlewares.JWTMiddleware())
	app.Use(cors.New(config))

	api.ApplyRoutes(app)

	// Listen and Server in 0.0.0.0:8080
	_ = app.Run(":8080")
}

