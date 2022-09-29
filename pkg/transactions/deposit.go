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
	Quantity  uint64
	Escrow    string
}

func NewDeposit(d DepositOptions, db *gorm.DB) (*Deposit, error) {
	err := balance.AddItems(d.AccountID, d.Item, d.Quantity, db)
	if err != nil {
		return nil, err
	}
	databaseDeposit := &database.Deposit{
		AccountID: d.AccountID,
		Item:      d.Item,
		Quantity:  d.Quantity,
		Escrow:    d.Escrow,
	}
	err = db.Create(databaseDeposit).Error
	if err != nil {
		return nil, err
	}
	return &Deposit{
		*databaseDeposit,
	}, nil
}
