package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JSON(c *gin.Context, message string, status int, data interface{}, errs []string) {
	responsedata := gin.H{
		"message": message,
		"data":    data,
		"errors":  errs,
		"status":  http.StatusText(status),
	}

	c.JSON(status, responsedata)
}
