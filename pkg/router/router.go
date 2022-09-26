package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/humbertovnavarro/farwater-bank/pkg/db"
)

func New() *gin.Engine {
	r := gin.Default()
	db := db.New()
	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})
	return r
}

func BindJSONOrWriteError(c *gin.Context, obj any) error {
	if err := c.BindJSON(obj); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return err
	}
	return nil
}
