package genre

import "github.com/gin-gonic/gin"

// ApplyRoutes apply router to gin router group
func ApplyRoutes(routerGroup *gin.RouterGroup) {
	genreRouter := routerGroup.Group("/genre")
	{
		genreRouter.POST("/", create)
		genreRouter.GET("/", getAll)
		genreRouter.DELETE("/:genre_id/soft", softDelete)
		genreRouter.DELETE("/:genre_id", permanentDelete)
	}
}
