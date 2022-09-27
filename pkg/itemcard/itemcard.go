package itemcard

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/humbertovnavarro/farwater-bank/pkg/token"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

var peppers []string

func init() {
	peppersString := os.Getenv("PEPPERS")
	peppers = strings.Split(peppersString, ",")
	if len(peppers) < 1 {
		logrus.Panicf("failed to enumerate peppers, is PEPPERS env set?")
	}
}

type ItemCard struct {
	gorm.Model
	AccountID uint
	Frozen    bool `gorm:"default:false"`
	Token     string
	Salt      string
	Pin       []byte
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
	pepper := peppers[rand.Int()%len(peppers)]
	if pepper == "" {
		logrus.Panic("empty pepper")
	}
	hashedPin, err := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("%s%s%s", pin, salt, pepper)), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	card := &ItemCard{
		AccountID: accountID,
		Frozen:    false,
		Token:     token,
		Salt:      salt,
		Pin:       hashedPin,
	}

	if err := db.Create(card).Error; err != nil {
		return nil, err
	}

	return card, nil
}

func ValidatePin(id uint, pin string, db *gorm.DB) error {
	card, err := Get(id, db)
	if err != nil {
		return err
	}
	for _, pepper := range peppers {
		password := []byte(fmt.Sprintf("%s%s%s", pin, card.Salt, pepper))
		if err := bcrypt.CompareHashAndPassword(card.Pin, password); err != nil {
			return nil
		}
	}
	return errors.New("match not found")
}

func IsFrozen(id uint, db *gorm.DB) bool {
	card, err := Get(id, db)
	if err != nil {
		logrus.Errorf("error while checking if itemcard is frozen: \n%s", err)
		return true
	}
	return card.Frozen
}

func Freeze(id uint, db *gorm.DB) error {
	return db.Model(&ItemCard{}).Where("id=?", id).Update("frozen", true).Error
}

func UnFreeze(id uint, db *gorm.DB) error {
	return db.Model(&ItemCard{}).Where("id=?", id).Update("frozen", false).Error
}
