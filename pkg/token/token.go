package token

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/humbertovnavarro/farwater-bank/pkg/config"
)

var userSecret []byte
var itemCardSecret []byte
var adminSecret []byte

var secrets [][]byte = make([][]byte, 0)

type TokenType = int

const (
	UserToken     TokenType = iota
	ItemCardToken TokenType = iota
	AdminToken    TokenType = iota
)

func init() {
	userSecret = []byte(config.AssertEnv("USER_SECRET"))
	secrets = append(secrets, userSecret)
	itemCardSecret = []byte(config.AssertEnv("ITEM_CARD_SECRET"))
	secrets = append(secrets, itemCardSecret)
	adminSecret = []byte(config.AssertEnv("ADMIN_SECRET"))
	secrets = append(secrets, adminSecret)
}

func SignedString(tokenType TokenType, subject string) (string, error) {
	secret := secrets[tokenType]
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": subject,
	})
	return token.SignedString(secret)
}

func ParseToken(token string, tokenType TokenType) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secrets[tokenType], nil
	})
}
