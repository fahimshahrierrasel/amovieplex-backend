package genre

import "github.com/gin-gonic/gin"

// ApplyRoutes apply router to gin router group
func ApplyRoutes(routerGroup *gin.RouterGroup) {
	ratings := routerGroup.Group("/genre")
	{
		ratings.POST("/", create)
	}
}
