package transactions

import (
	"testing"

	"github.com/humbertovnavarro/farwater-bank/pkg/account"
	"github.com/humbertovnavarro/farwater-bank/pkg/balance"
	"github.com/humbertovnavarro/farwater-bank/pkg/database"
	"github.com/humbertovnavarro/farwater-bank/pkg/mocks"
	"github.com/stretchr/testify/assert"
)

func assertBalance(t *testing.T, b *balance.Balance, quantity uint64) {
	if !assert.NotNil(t, b) {
		return
	}
	assert.Equal(t, "minecraft:dirt", b.Item)
	assert.Equal(t, quantity, b.Quantity)
}

func TestDeposit(t *testing.T) {
	db := mocks.NewMockDB()
	db.Create(&database.Account{})
	a, err := account.GetByID(1, db)
	if !assert.Nil(t, err) {
		return
	}
	if !assert.NotNil(t, a) {
		return
	}

	for i := 0; i < 4; i++ {
		NewDeposit(DepositOptions{
			AccountID: a.ID,
			Item:      "minecraft:dirt",
			Quantity:  64,
			Escrow:    "foo",
		}, db)
	}
	var deposits []database.Deposit
	err = db.Model(&database.Deposit{}).Where("account_id = ?", a.ID).Find(&deposits).Error
	assert.Len(t, deposits, 4)
	assert.Nil(t, err)

	b, err := balance.Get(a.ID, "minecraft:dirt", db)
	assert.Nil(t, err)
	assertBalance(t, b, 64*4)
}
