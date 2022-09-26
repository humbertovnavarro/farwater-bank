package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/humbertovnavarro/farwater-bank/pkg/middleware"
	"github.com/humbertovnavarro/farwater-bank/pkg/token"
)

type Registration struct {
	MinecraftName string
	Escrow        string
}

func Register(c *gin.Context) {
	authorization := c.MustGet("authorization").(*middleware.RequestAuthorization)
	if authorization.TokenType != token.AdminToken {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}
	registration := &Registration{}
	if err := c.BindJSON(registration); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}
