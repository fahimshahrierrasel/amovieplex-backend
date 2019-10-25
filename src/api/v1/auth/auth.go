package auth

import "github.com/gin-gonic/gin"

// ApplyRoutes apply router to gin router group
func ApplyRoutes(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", register)
		auth.POST("/login", login)
	}
}
