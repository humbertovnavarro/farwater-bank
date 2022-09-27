package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/humbertovnavarro/farwater-bank/pkg/account"
	"github.com/humbertovnavarro/farwater-bank/pkg/middleware"
	"github.com/humbertovnavarro/farwater-bank/pkg/token"
	"gorm.io/gorm"
)

type RegistrationRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Pin      string `json:"pin"`
}

func Register(c *gin.Context) {
	authorization := c.MustGet("authorization").(*middleware.RequestAuthorization)
	if authorization.TokenType != token.AdminToken {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}
	registration := &RegistrationRequest{}
	if err := c.BindJSON(registration); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	db := c.MustGet("db").(*gorm.DB)
	_, err := account.Register(registration.Username, registration.Password, registration.Pin, db)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "could not register account",
		})
	}
}
