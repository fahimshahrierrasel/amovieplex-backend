package rating

import (
	"amovieplex-backend/src/middlewares"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes apply router to gin router group
func ApplyRoutes(routerGroup *gin.RouterGroup) {
	ratings := routerGroup.Group("/rating")
	{
		ratings.POST("/", middlewares.Authorized, create)
		ratings.GET("/", getAll)
		ratings.GET("/:rating_id", getRating)
		ratings.PUT("/:rating_id", middlewares.Authorized, updateRating)
		ratings.DELETE("/:rating_id/soft", middlewares.Authorized, softDelete)
		ratings.DELETE("/:rating_id", middlewares.Authorized, permanentDelete)
	}
}
