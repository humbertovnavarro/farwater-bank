package atm

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/humbertovnavarro/farwater-bank/pkg/account"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RegistrationRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Pin      string `json:"pin"`
}

func Register(c *gin.Context) {
	registration := &RegistrationRequest{}
	if err := c.BindJSON(registration); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	db := c.MustGet("db").(*gorm.DB)
	a, err := account.Register(registration.Username, registration.Password, registration.Pin, db)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})
		logrus.Error(err)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"minecraft_uuid": a.MinecraftUUID,
		"account_id":     a.ID,
	})
}
