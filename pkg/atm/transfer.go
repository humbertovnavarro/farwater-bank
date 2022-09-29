package atm

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/humbertovnavarro/farwater-bank/pkg/token"
	"github.com/humbertovnavarro/farwater-bank/pkg/transactions"
	"gorm.io/gorm"
)

type TransferRequest struct {
	AccountID   uint   `json:"account_id"`
	ToAccountID uint   `json:"to_account_id"`
	Item        string `json:"item"`
	Quantity    uint64 `json:"quantity"`
	Escrow      string `json:"escrow"`
}

func Transfer(c *gin.Context) {
	auth := c.MustGet("authorization").(*token.Token)
	request := &TransferRequest{}
	if err := c.BindJSON(request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "bad JSON syntax",
		})
		return
	}
	db := c.MustGet("db").(*gorm.DB)
	w, err := transactions.NewTransfer(transactions.TransferOptions{
		AccountID:   request.AccountID,
		Item:        request.Item,
		Quantity:    request.Quantity,
		Escrow:      auth.Subject,
		ToAccountID: request.ToAccountID,
	}, db)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, w)
}
