package rating

import "github.com/gin-gonic/gin"

func ApplyRoutes(routerGroup *gin.RouterGroup) {
	ratings := routerGroup.Group("/rating")
	{
		ratings.POST("/", create)
		ratings.GET("/all", getAll)
	}
}