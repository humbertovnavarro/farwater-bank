package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindJSONOrWriteError(c *gin.Context, obj any) error {
	if err := c.BindJSON(obj); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return err
	}
	return nil
}
