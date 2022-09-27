package transactions

import (
	"github.com/humbertovnavarro/farwater-bank/pkg/balance"
	"github.com/humbertovnavarro/farwater-bank/pkg/database"
	"gorm.io/gorm"
)

type Transfer struct {
	database.Transfer
}

type TransferOptions struct {
	AccountID   uint
	Item        string
	Amount      uint64
	ToAccountID uint
	Escrow      string
}

func NewTransfer(t TransferOptions, db *gorm.DB) (transaction *Transfer, from *balance.Balance, to *balance.Balance, err error) {
	tx := &Transfer{
		database.Transfer{
			Item:      t.Item,
			AccountID: t.AccountID,
			Amount:    t.Amount,
			Escrow:    t.Escrow,
		},
	}
	fromB, err := balance.RemoveItems(t.AccountID, t.Item, t.Amount, db)
	if err != nil {
		return nil, nil, nil, err
	}
	toB, err := balance.AddItems(t.ToAccountID, t.Item, t.Amount, db)
	if err != nil {
		// reverse first transaction
		balance.AddItems(t.AccountID, t.Item, t.Amount, db)
		return nil, nil, nil, err
	}
	if err := db.Create(tx).Error; err != nil {
		return nil, nil, nil, err
	}
	return tx, fromB, toB, nil
}
