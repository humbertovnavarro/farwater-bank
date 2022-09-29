package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/humbertovnavarro/farwater-bank/pkg/atm"
	"github.com/humbertovnavarro/farwater-bank/pkg/database"
	"github.com/humbertovnavarro/farwater-bank/pkg/middleware"
)

func New() *gin.Engine {
	r := gin.Default()
	db := database.New()
	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})
	r.GET("/atm/account/:uuid", middleware.ATMAuthentication, atm.Account)
	r.POST("/atm/transfer", middleware.ATMAuthentication, atm.Transfer)
	r.POST("/atm/deposit", middleware.ATMAuthentication, atm.Deposit)
	r.POST("/atm/balance", middleware.ATMAuthentication, atm.Balance)
	r.POST("/atm/withdrawal", middleware.ATMAuthentication, atm.Withdrawal)
	r.POST("/atm/register", middleware.ATMAuthentication, atm.Register)
	r.POST("/atm/verify-pin", middleware.ATMAuthentication, atm.VerifyPin)
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
