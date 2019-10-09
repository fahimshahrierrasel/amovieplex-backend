package helpers

import "github.com/gin-gonic/gin"

// MakeResponse make json response from data, error and message
func MakeResponse(data interface{}, error bool, message string) gin.H {
	return gin.H{
		"error":   error,
		"data":    data,
		"message": message,
	}
}
