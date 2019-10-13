package movie

import "github.com/gin-gonic/gin"

func ApplyRoutes(routerGroup *gin.RouterGroup) {
	movies := routerGroup.Group("/movies")
	{
		movies.POST("/", create)
	}
}