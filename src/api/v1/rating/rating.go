package rating

import "github.com/gin-gonic/gin"

// ApplyRoutes apply router to gin router group
func ApplyRoutes(routerGroup *gin.RouterGroup) {
	ratings := routerGroup.Group("/rating")
	{
		ratings.POST("/", create)
		ratings.GET("/", getAll)
		ratings.DELETE("/:rating_id/soft", softDelete)
		ratings.DELETE("/:rating_id", permanentDelete)
	}
}
