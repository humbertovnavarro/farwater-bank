package atm

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/humbertovnavarro/farwater-bank/pkg/account"
	"github.com/humbertovnavarro/farwater-bank/pkg/token"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type VerifyPinRequest struct {
	Pin           string `json:"pin"`
	MinecraftUUID string `json:"minecraft_uuid"`
}

func VerifyPin(c *gin.Context) {
	authorization := c.MustGet("authorization").(*token.Token)
	if !(authorization.Type == token.ATMToken) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "bad token",
		})
		logrus.Panic("got wrong token type while trying to verify a pin")
		return
	}
	request := &VerifyPinRequest{}
	if err := c.BindJSON(request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
	}
	db := c.MustGet("db").(*gorm.DB)
	a, err := account.GetByUUID(request.MinecraftUUID, db)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "account not found",
		})
	}
	err = a.VerifyPin(request.Pin)
	if err == nil {
		c.AbortWithStatus(http.StatusOK)
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
	}
}
