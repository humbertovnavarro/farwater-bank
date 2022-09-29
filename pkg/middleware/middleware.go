package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/humbertovnavarro/farwater-bank/pkg/token"
	"github.com/sirupsen/logrus"
)

var UserAuthentication = authentication(token.UserToken)
var ATMAuthentication = authentication(token.ATMToken)

func writeUnauthorizedError(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error": "unauthorized",
	})
}

func authentication(acceptableTokenTypes ...token.TokenType) gin.HandlerFunc {
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
		token, err := token.ParseToken(tokenString)
		if err != nil {
			writeUnauthorizedError(c)
			return
		}
		for _, tokenType := range acceptableTokenTypes {
			if tokenType == token.Type {
				c.Set("authorization", token)
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
	}
}
