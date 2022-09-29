package atm

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/humbertovnavarro/farwater-bank/pkg/account"
	"gorm.io/gorm"
)

func Account(c *gin.Context) {
	uuid := c.Param("uuid")
	db := c.MustGet("db").(*gorm.DB)
	a, err := account.GetByUUID(uuid, db)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "not found",
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"id":     a.ID,
		"frozen": a.Frozen,
	})
}
