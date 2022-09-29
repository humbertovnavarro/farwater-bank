package token

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

var userSecret []byte
var ATMSecret []byte

var secrets [][]byte = make([][]byte, 0)
var headers []string = []string{"u", "a"}
var headerMap map[string]TokenType = make(map[string]int)

type TokenType = int

type Token struct {
	Type    TokenType
	Subject string
}

const (
	UserToken TokenType = iota
	ATMToken  TokenType = iota
)

func init() {
	userSecret = []byte(os.Getenv("USER_SECRET"))
	secrets = append(secrets, userSecret)
	ATMSecret = []byte(os.Getenv("ATM_SECRET"))
	secrets = append(secrets, ATMSecret)
	headerMap["u"] = UserToken
	headerMap["a"] = ATMToken
}

func SignedString(tokenType TokenType, subject string) (string, error) {
	secret := secrets[tokenType]
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": subject,
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", headers[tokenType], tokenString), nil
}

func ParseToken(rawTokenString string) (*Token, error) {
	if len(rawTokenString) < 1 {
		return nil, errors.New("bad token")
	}
	tokenString := rawTokenString[1:]
	tokenType := headerMap[rawTokenString[0:1]]
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
