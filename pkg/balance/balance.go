package balance

import (
	"fmt"

	"github.com/humbertovnavarro/farwater-bank/pkg/database"
	"gorm.io/gorm"
)

type Balance struct {
	database.Balance
}

func AddItems(accountID uint, item string, amount uint64, db *gorm.DB) (*Balance, error) {
	balance := &Balance{}
	err := db.Find(balance).Where("account_id = ? AND item = ?").Error

	if err == gorm.ErrRecordNotFound {
		newBalance := &Balance{
			database.Balance{
				AccountID: accountID,
				Item:      item,
				Quantity:  amount,
			},
		}
		err := db.Create(newBalance).Error
		if err != nil {
			return nil, err
		}
		return newBalance, nil
	}

	err = db.Model(&Balance{}).Where("id = ?", balance.ID).Update("quantity", balance.Quantity+amount).Error
	if err != nil {
		return nil, err
	}
	return &Balance{
		database.Balance{
			AccountID: accountID,
			Item:      item,
			Quantity:  balance.Quantity + amount,
		},
	}, nil
}

func RemoveItems(accountID uint, item string, amount uint64, db *gorm.DB) (*Balance, error) {
	balance := &Balance{}
	err := db.Find(balance).Where("account_id = ? AND item = ?").Error

	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("no balance found satisfying account: %d, item: %s", accountID, item)
	}
	if balance.Quantity < amount {
		return nil, fmt.Errorf("not enough items to cover amount: %d, have %d", amount, balance.Quantity)
	}
	err = db.Model(&Balance{}).Where("id = ?", balance.ID).Update("quantity", balance.Quantity+amount).Error
	if err != nil {
		return nil, err
	}
	return &Balance{
		database.Balance{
			AccountID: accountID,
			Item:      item,
			Quantity:  balance.Quantity + amount,
		},
	}, nil
}
