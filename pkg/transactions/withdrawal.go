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
	item      string
	Amount    uint64
	Escrow    string
}

func NewWithdrawal(w WithdrawalOptions, db *gorm.DB) (*Withdrawal, *balance.Balance, error) {
	tx := &Withdrawal{
		database.Withdrawal{
			Item:      w.item,
			AccountID: w.AccountID,
			Amount:    w.Amount,
			Escrow:    w.Escrow,
		},
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
