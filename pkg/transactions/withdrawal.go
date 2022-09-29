package transactions

import (
	"github.com/humbertovnavarro/farwater-bank/pkg/balance"
	"github.com/humbertovnavarro/farwater-bank/pkg/database"
	"gorm.io/gorm"
)

type Withdrawal struct {
	database.Withdrawal
}

type WithdrawalOptions struct {
	AccountID uint
	Item      string
	Quantity  uint64
	Escrow    string
}

func NewWithdrawal(w WithdrawalOptions, db *gorm.DB) (*Withdrawal, error) {
	err := balance.RemoveItems(w.AccountID, w.Item, w.Quantity, db)
	if err != nil {
		return nil, err
	}
	databaseWithdrawal := &database.Withdrawal{
		AccountID: w.AccountID,
		Item:      w.Item,
		Quantity:  w.Quantity,
		Escrow:    w.Escrow,
	}
	err = db.Create(databaseWithdrawal).Error
	if err != nil {
		return nil, err
	}
	return &Withdrawal{
		*databaseWithdrawal,
	}, nil
}
