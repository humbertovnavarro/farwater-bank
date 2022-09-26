package itemcard

import (
	"github.com/humbertovnavarro/farwater-bank/pkg/account"
	"github.com/humbertovnavarro/farwater-bank/pkg/token"
	"github.com/sirupsen/logrus"

	"gorm.io/gorm"
)

type ItemCard struct {
	gorm.Model
	AccountID uint
	Frozen    bool `gorm:"default:false"`
	Token     string
}

func Get(u *account.Account, db *gorm.DB) (*ItemCard, error) {
	card := &ItemCard{}
	if err := db.First(card, u.ID).Error; err != nil {
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

func Issue(accountID uint, minecraftUUID string, db *gorm.DB) (*ItemCard, error) {
	Delete(accountID, db)
	token, err := token.SignedString(token.ItemCardToken, minecraftUUID)
	if err != nil {
		return nil, err
	}
	card := &ItemCard{
		AccountID: accountID,
		Frozen:    false,
		Token:     token,
	}
	if err := db.Create(card).Error; err != nil {
		return nil, err
	}
	return card, nil
}

func IsFrozen(u *account.Account, db *gorm.DB) bool {
	card, err := Get(u, db)
	if err != nil {
		logrus.Errorf("error while checking if itemcard is frozen: \n%s", err)
		return true
	}
	return card.Frozen
}
