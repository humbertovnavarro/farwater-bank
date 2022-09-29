package token

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

var userSecret []byte
var adminSecret []byte

var secrets [][]byte = make([][]byte, 0)

type TokenType = int

type Token struct {
	Type    TokenType
	Subject string
}

const (
	UserToken  TokenType = iota
	AdminToken TokenType = iota
)

func init() {
	userSecret = []byte(os.Getenv("USER_SECRET"))
	secrets = append(secrets, userSecret)
	adminSecret = []byte(os.Getenv("ADMIN_SECRET"))
	secrets = append(secrets, adminSecret)
}

func SignedString(tokenType TokenType, subject string) (string, error) {
	secret := secrets[tokenType]
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": subject,
	})
	return token.SignedString(secret)
}

func ParseToken(tokenString string, tokenType TokenType) (*Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secrets[tokenType], nil
	})
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(jwt.MapClaims)
	if claims["sub"] == nil {
		return nil, errors.New("nil subject on token")
	}
	subject := claims["sub"].(string)
	if subject == "" {
		return nil, errors.New("empty subject on token")
	}
	return &Token{
		Type:    tokenType,
		Subject: subject,
	}, nil
}
