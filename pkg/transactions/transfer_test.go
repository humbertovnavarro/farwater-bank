package transactions

import (
	"testing"

	"github.com/humbertovnavarro/farwater-bank/pkg/balance"
	"github.com/humbertovnavarro/farwater-bank/pkg/database"
	mocks_test "github.com/humbertovnavarro/farwater-bank/pkg/mocks"
	"github.com/stretchr/testify/assert"
)

func TestTransfer(t *testing.T) {
	db := mocks_test.NewMockDB()
	db.Create(&database.Account{})
	db.Create(&database.Account{})
	var accountA uint = 1
	var accountB uint = 2
	balance.AddItems(accountA, "minecraft:dirt", 64, db)
	balance.AddItems(accountB, "minecraft:dirt", 64, db)
	tr, err := NewTransfer(TransferOptions{
		AccountID:   accountA,
		ToAccountID: accountB,
		Item:        "minecraft:dirt",
		Quantity:    32,
	}, db)
	assert.Nil(t, err)
	assert.NotNil(t, tr)
	balanceA, err := balance.Get(accountA, "minecraft:dirt", db)
	assert.Nil(t, err)
	balanceB, err := balance.Get(accountB, "minecraft:dirt", db)
	assert.Nil(t, err)
	assertBalance(t, balanceA, 32)
	assertBalance(t, balanceB, 64+32)
	dbTx := &database.Transfer{}
	err = db.Model(&database.Transfer{}).Where("id = ?", tr.ID).First(dbTx).Error
	assert.Nil(t, err)
	assert.Equal(t, tr.ID, dbTx.ID)
	assert.Equal(t, tr.Quantity, dbTx.Quantity)
	assert.Equal(t, tr.Item, dbTx.Item)
}
