package movie

import (
	"amovieplex-backend/src/middlewares"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(routerGroup *gin.RouterGroup) {
	movies := routerGroup.Group("/movies")
	{
		movies.POST("/", middlewares.Authorized, create)
		movies.GET("/", getAll)
		movies.GET("/:movie_id", getMovie)
		movies.PUT("/:movie_id", middlewares.Authorized, updateMovie)
		movies.DELETE("/:movie_id/soft", middlewares.Authorized, softDelete)
		movies.DELETE("/:movie_id", middlewares.Authorized, permanentDelete)
	}
}
