package balance

import (
	"errors"
	"fmt"

	"github.com/humbertovnavarro/farwater-bank/pkg/account"
	"github.com/humbertovnavarro/farwater-bank/pkg/database"
	"gorm.io/gorm"
)

type Balance struct {
	database.Balance
}

func Get(accountID uint, item string, db *gorm.DB) (*Balance, error) {
	b := &database.Balance{}
	err := db.Model(&database.Balance{}).Where("id = ? AND item = ?", accountID, item).First(b).Error
	if err != nil {
		return nil, err
	}
	return &Balance{
		*b,
	}, nil
}

func New(accountID uint, item string, startingQuantity uint64, db *gorm.DB) (*Balance, error) {
	_, err := account.GetByID(accountID, db)
	if err != nil {
		return nil, errors.New("account does not exist")
	}
	b := &database.Balance{
		AccountID: accountID,
		Item:      item,
		Quantity:  startingQuantity,
	}
	err = db.Create(b).Error
	if err != nil {
		return nil, err
	}
	return &Balance{
		*b,
	}, nil
}

func AddItems(accountID uint, item string, quantity uint64, db *gorm.DB) error {
	tx := db.Model(&database.Balance{}).Where("account_id = ? AND item = ?", accountID, item).Update("quantity", gorm.Expr("quantity + ?", quantity))
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		_, err := New(accountID, item, quantity, db)
		return err
	}
	return nil
}

func RemoveItems(accountID uint, item string, quantity uint64, db *gorm.DB) error {
	existing, err := Get(accountID, item, db)
	if err != nil {
		return err
	}
	if existing.Quantity < quantity {
		return fmt.Errorf("%d attempted to overdraft %s", accountID, item)
	}
	tx := db.Model(&database.Balance{}).Where("account_id = ? AND item = ?", accountID, item).Update("quantity", gorm.Expr("quantity - ?", quantity))
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		_, err := New(accountID, item, quantity, db)
		return err
	}
	return nil
}
