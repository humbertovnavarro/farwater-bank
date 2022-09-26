package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

type RequestAuthorization struct {
	Subject   string
	TokenType token.TokenType
}

func authentication(tokenType token.TokenType) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header["Authentication"]
		if authHeader[1] == "" {
			writeUnauthorizedError(c)
			return
		}
		tokenString := authHeader[0]
		if len(tokenString) < 10 {
			writeUnauthorizedError(c)
			return
		}
		verifiedToken, err := token.ParseToken(tokenString, tokenType)
		if err != nil {
			logrus.Error(err)
			writeUnauthorizedError(c)
			return
		}
		claims := verifiedToken.Claims.(jwt.StandardClaims)
		authorization := &RequestAuthorization{
			Subject:   claims.Subject,
			TokenType: tokenType,
		}
		c.Set("authorization", authorization)
	}
}
