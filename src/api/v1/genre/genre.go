package genre

import (
	"amovieplex-backend/src/middlewares"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes apply router to gin router group
func ApplyRoutes(routerGroup *gin.RouterGroup) {
	genreRouter := routerGroup.Group("/genres")
	{
		genreRouter.POST("", middlewares.Authorized, create)
		genreRouter.GET("", getAll)
		genreRouter.GET("/:genre_id", getGenre)
		genreRouter.PUT("/:genre_id", middlewares.Authorized, updateGenre)
		genreRouter.DELETE("/:genre_id/soft", middlewares.Authorized, softDelete)
		genreRouter.DELETE("/:genre_id", middlewares.Authorized, permanentDelete)
	}
}
