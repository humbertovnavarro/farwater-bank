package account

import (
	"os"
	"strings"

	"github.com/humbertovnavarro/farwater-bank/pkg/database"
	"github.com/humbertovnavarro/farwater-bank/pkg/minecraft"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Account struct {
	database.Account
}

var peppers []string

func init() {
	peppersString := os.Getenv("PEPPERS")
	peppers = strings.Split(peppersString, ",")
	if len(peppers) < 1 {
		logrus.Panicf("failed to enumerate peppers, is PEPPERS env set?")
	}
}

func GetByID(id uint, db *gorm.DB) (*Account, error) {
	account := &Account{}
	if err := db.First(account, id).Error; err != nil {
		return nil, err
	}
	return account, nil
}

func GetByUUID(uuid string, db *gorm.DB) (*Account, error) {
	account := &Account{}
	if err := db.First(account, "minecraft_uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return account, nil
}

func Register(username string, password string, pin string, db *gorm.DB) (*Account, error) {
	uuid, err := minecraft.FetchUUID(username)
	if err != nil {
		return nil, err
	}
	passwordSalt, hashedPassword, err := HashSecret(password)
	if err != nil {
		return nil, err
	}
	pinSalt, hashedPin, err := HashSecret(pin)
	account := &Account{
		database.Account{
			MinecraftUUID: uuid,
			Password:      hashedPassword,
			PasswordSalt:  passwordSalt,
			Pin:           hashedPin,
			PinSalt:       pinSalt,
		},
	}
	if err := db.FirstOrCreate(account, "minecraft_uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	if err != nil {
		logrus.Error(err)
	}
	return account, nil
}

func (a *Account) VerifyPin(pin string) error {
	return VerifySecret(pin, a.Pin, a.PinSalt)
}

func (a *Account) VerifyPassword(password string) error {
	return VerifySecret(password, a.Password, a.PasswordSalt)
}
