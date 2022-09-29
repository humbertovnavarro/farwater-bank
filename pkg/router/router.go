package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/humbertovnavarro/farwater-bank/pkg/database"
	"github.com/humbertovnavarro/farwater-bank/pkg/middleware"
	"github.com/humbertovnavarro/farwater-bank/pkg/routes"
)

func New() *gin.Engine {
	r := gin.Default()
	db := database.New()
	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})
	r.POST("/atm/withdrawal", middleware.AdminAuthentication, routes.Withdrawal)
	r.POST("/atm/register", middleware.AdminAuthentication, routes.Register)
	r.POST("/atm/verify-pin", middleware.AdminAuthentication, routes.VerifyPin)
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
