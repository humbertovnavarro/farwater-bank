package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/humbertovnavarro/farwater-bank/pkg/token"
	"github.com/humbertovnavarro/farwater-bank/pkg/transactions"
	"gorm.io/gorm"
)

type DepositRequest struct {
	AccountID uint   `json:"account_id"`
	Item      string `json:"item"`
	Quantity  uint64 `json:"quantity"`
}

func Deposit(c *gin.Context) {
	auth := c.MustGet("authorization").(*token.Token)
	request := &DepositRequest{}
	if err := c.BindJSON(request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "bad JSON syntax",
		})
		return
	}
	db := c.MustGet("db").(*gorm.DB)
	d, err := transactions.NewDeposit(transactions.DepositOptions{
		AccountID: request.AccountID,
		Item:      request.Item,
		Quantity:  request.Quantity,
		Escrow:    auth.Subject,
	}, db)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, d)
}
