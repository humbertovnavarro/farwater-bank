package itemcard

import (
	"crypto/aes"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/humbertovnavarro/farwater-bank/pkg/token"
	"github.com/sirupsen/logrus"

	"gorm.io/gorm"
)

const ENCRYPTION_TAG = "_itemcard_"

var peppers []string

func init() {
	peppersString := os.Getenv("PEPPERS")
	peppers = strings.Split(peppersString, ",")
}

type ItemCard struct {
	gorm.Model
	AccountID uint
	Frozen    bool `gorm:"default:false"`
	Token     string
	Salt      string
}

func Get(id uint, db *gorm.DB) (*ItemCard, error) {
	card := &ItemCard{}
	if err := db.First(card, id).Error; err != nil {
		return nil, err
	}
	return card, nil
}

func Delete(accountID uint, db *gorm.DB) error {
	if err := db.Where("account_id=?", accountID).Delete(&ItemCard{}).Error; err != nil {
		return err
	}
	return nil
}

func Issue(accountID uint, pin string, db *gorm.DB) (*ItemCard, error) {
	Delete(accountID, db)
	token, err := token.SignedString(token.ItemCardToken, fmt.Sprintf("%d", accountID))
	if err != nil {
		return nil, err
	}
	salt := uuid.New().String()
	encryptedToken, err := encryptToken(token, pin, salt)
	if err != nil {
		return nil, err
	}
	card := &ItemCard{
		AccountID: accountID,
		Frozen:    false,
		Token:     encryptedToken,
		Salt:      salt,
	}
	if err := db.Create(card).Error; err != nil {
		return nil, err
	}
	return card, nil
}

func IsFrozen(id uint, db *gorm.DB) bool {
	card, err := Get(id, db)
	if err != nil {
		logrus.Errorf("error while checking if itemcard is frozen: \n%s", err)
		return true
	}
	return card.Frozen
}

func encryptToken(token string, pin string, salt string) (string, error) {
	pepper := peppers[rand.Intn(len(peppers))]
	key := fmt.Sprintf("%s%s%s", pin, salt, pepper)
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	out := make([]byte, len(token))
	c.Encrypt(out, []byte(token))
	return hex.EncodeToString(out), nil
}

func Decrypt(id uint, encryptedToken string, pin string, db *gorm.DB) (string, *ItemCard, error) {
	card, err := Get(id, db)
	if err != nil {
		return "", nil, err
	}
	salt := card.Salt
	for _, pepper := range peppers {
		key := []byte(fmt.Sprintf("%s%s%s", pin, salt, pepper))
		ciphertext, _ := hex.DecodeString(encryptedToken)
		c, err := aes.NewCipher(key)
		if err != nil {
			continue
		}
		out := make([]byte, len(ciphertext))
		c.Decrypt(out, ciphertext)
		decrypted := string(out)
		if strings.HasPrefix(decrypted, ENCRYPTION_TAG) {
			return decrypted, card, nil
		}
	}
	return "", nil, errors.New("wrong pin")
}
