package api

import (
	"amovieplex-backend/src/api/v1"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Engine
func ApplyRoutes(engine *gin.Engine) {
	api := engine.Group("/api")
	{
		v1.ApplyRoutes(api)
	}
}
