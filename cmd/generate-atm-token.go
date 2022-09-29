package main

import (
	"os"

	"github.com/humbertovnavarro/farwater-bank/pkg/token"
)

func main() {
	if len(os.Args) != 3 {
		os.Stderr.WriteString("usage: atm_id atm_secret\n")
		return
	}
	subject := os.Args[1]
	secret := os.Args[2]
	if subject == "" {
		os.Stderr.WriteString("you must provide an atm id\n")
		return
	}
	if secret == "" {
		os.Stderr.WriteString("you must provide the atm secret\n")
		return
	}
	token.ATMSecret = []byte(secret)
	tokenString, err := token.SignedString(token.ATMToken, subject)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		return
	}
	os.Stdout.WriteString(tokenString + "\n")
}
