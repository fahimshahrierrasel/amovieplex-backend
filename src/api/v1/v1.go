package v1

import (
	"amovieplex-backend/src/api/helpers"
	"amovieplex-backend/src/api/v1/rating"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ping checks if the api is accessible
func ping(c *gin.Context) {
	data := map[string]interface{}{
		"status": "pong",
		"number": 1,
	}
	c.JSON(http.StatusOK, helpers.MakeResponse(data, false, ""))
}

// ApplyRoutes applies router to gin Router
func ApplyRoutes(routerGroup *gin.RouterGroup) {
	v1 := routerGroup.Group("/v1")
	{
		v1.GET("/ping", ping)
		rating.ApplyRoutes(v1)
	}
}
