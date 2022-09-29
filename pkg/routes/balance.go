package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/humbertovnavarro/farwater-bank/pkg/balance"
	"gorm.io/gorm"
)

type BalanceRequest struct {
	AccountID uint   `json:"account_id"`
	Item      string `json:"item"`
}

func Balance(c *gin.Context) {
	req := &BalanceRequest{}
	if err := c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "bad JSON",
		})
	}
	db := c.MustGet("db").(*gorm.DB)
	b, err := balance.Get(req.AccountID, req.Item, db)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"quantity": 0,
		})
	} else {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"quantity": b.Quantity,
		})
	}
}
