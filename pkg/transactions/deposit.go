package transactions

import (
	"github.com/humbertovnavarro/farwater-bank/pkg/balance"
	"github.com/humbertovnavarro/farwater-bank/pkg/database"
	"gorm.io/gorm"
)

type Deposit struct {
	database.Deposit
}

type DepositOptions struct {
	AccountID uint
	Item      string
	Amount    uint64
	Escrow    string
}

func NewDeposit(d DepositOptions, db *gorm.DB) (*Deposit, *balance.Balance, error) {
	tx := &Deposit{
		database.Deposit{
			Item:      d.Item,
			AccountID: d.AccountID,
			Amount:    d.Amount,
			Escrow:    d.Escrow,
		},
	}
	b, err := balance.AddItems(d.AccountID, d.Item, d.Amount, db)
	if err != nil {
		return nil, nil, err
	}
	if err := db.Create(tx).Error; err != nil {
		return nil, nil, err
	}
	return tx, b, nil
}
