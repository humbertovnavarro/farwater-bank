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
	Quantity    uint64
	ToAccountID uint
	Escrow      string
}

func NewTransfer(t TransferOptions, db *gorm.DB) (*Transfer, error) {
	err := balance.RemoveItems(t.AccountID, t.Item, t.Quantity, db)
	if err != nil {
		return nil, err
	}
	err = balance.AddItems(t.ToAccountID, t.Item, t.Quantity, db)
	if err != nil {
		balance.AddItems(t.AccountID, t.Item, t.Quantity, db)
		return nil, err
	}
	databaseTransfer := &database.Transfer{
		AccountID:   t.AccountID,
		ToAccountID: t.ToAccountID,
		Item:        t.Item,
		Quantity:    t.Quantity,
		Escrow:      t.Escrow,
	}
	err = db.Create(databaseTransfer).Error
	if err != nil {
		return nil, err
	}
	return &Transfer{
		*databaseTransfer,
	}, nil
}
