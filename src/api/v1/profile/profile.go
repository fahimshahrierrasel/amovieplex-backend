package profile

import (
	"amovieplex-backend/src/middlewares"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes apply router to gin router group
func ApplyRoutes(router *gin.RouterGroup) {
	profile := router.Group("/profile")
	{
		profile.POST("", middlewares.Authorized, create)
		profile.GET("/:user_id", middlewares.Authorized, getProfileByUserID)
	}
}
