package account

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var peppers []string

func HashSecret(secret string) (secretSalt string, secretHex string, err error) {
	salt := uuid.New().String()
	pepper := peppers[rand.Int()%len(peppers)]
	passwordString := fmt.Sprintf("%s%s%s", secret, salt, pepper)
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(passwordString), bcrypt.DefaultCost)
	return salt, hex.EncodeToString(passwordHash), err
}

func VerifySecret(secretText string, comparisonHex string, salt string) error {
	hashedSecret, err := hex.DecodeString(comparisonHex)
	if err != nil {
		return err
	}
	for _, pepper := range peppers {
		secret := []byte(fmt.Sprintf("%s%s%s", secretText, salt, pepper))
		if err := bcrypt.CompareHashAndPassword(hashedSecret, secret); err == nil {
			return nil
		}
	}
	return errors.New("invalid secret")
}
