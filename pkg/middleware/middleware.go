package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/humbertovnavarro/farwater-bank/pkg/token"
	"github.com/sirupsen/logrus"
)

var UserAuthentication = authentication(token.UserToken)
var AdminAuthentication = authentication(token.AdminToken)

func writeUnauthorizedError(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error": "unauthorized",
	})
}

func authentication(tokenType token.TokenType) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header["Authorization"]
		if len(authHeader) < 1 || authHeader[0] == "" {
			writeUnauthorizedError(c)
			logrus.Errorf("authentication failed: %s", c.RemoteIP())
			return
		}
		tokenString := authHeader[0][strings.Index(authHeader[0], " ")+1:]
		if len(tokenString) < 10 {
			writeUnauthorizedError(c)
			logrus.Errorf("authentication failed: %s", c.RemoteIP())
			return
		}
		token, err := token.ParseToken(tokenString, tokenType)
		if err != nil {
			logrus.Error(err)
			writeUnauthorizedError(c)
			logrus.Errorf("authentication failed: %s", c.RemoteIP())
			return
		}
		if token.Type != tokenType {
			writeUnauthorizedError(c)
			logrus.Errorf("authentication failed: %s", c.RemoteIP())
		}
		c.Set("authorization", token)
	}
}
