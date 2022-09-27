package transactions

import (
	"github.com/humbertovnavarro/farwater-bank/pkg/balance"
	"gorm.io/gorm"
)

type withdrawal struct {
	gorm.Model
	Item      string
	Amount    uint64
	AccountID uint
	Escrow    string
}

type WithdrawalOptions struct {
	AccountID uint
	item      string
	Amount    uint64
	Escrow    string
}

func NewWithdrawal(w WithdrawalOptions, db *gorm.DB) (*withdrawal, *balance.Balance, error) {
	tx := &withdrawal{
		Item:      w.item,
		AccountID: w.AccountID,
		Amount:    w.Amount,
		Escrow:    w.Escrow,
	}
	b, err := balance.RemoveItems(w.AccountID, w.item, w.Amount, db)
	if err != nil {
		return nil, nil, err
	}
	if err := db.Create(tx).Error; err != nil {
		return nil, nil, err
	}
	return tx, b, nil
}
