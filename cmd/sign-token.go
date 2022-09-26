package main

import (
	"os"

	"github.com/golang-jwt/jwt"
)

func main() {
	if len(os.Args) != 3 {
		os.Stderr.WriteString("usage:  subect secret\n")
		return
	}
	subject := os.Args[1]
	secret := os.Args[2]
	if subject == "" {
		os.Stderr.WriteString("you must provide a bearer id\n")
		return
	}
	if secret == "" {
		os.Stderr.WriteString("you must provide a secret\n")
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": subject,
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		return
	}
	os.Stdout.WriteString(tokenString + "\n")
}
