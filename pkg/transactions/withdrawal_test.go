package transactions

import (
	"testing"

	"github.com/humbertovnavarro/farwater-bank/pkg/account"
	"github.com/humbertovnavarro/farwater-bank/pkg/balance"
	"github.com/humbertovnavarro/farwater-bank/pkg/database"
	mocks_test "github.com/humbertovnavarro/farwater-bank/pkg/mocks"
	"github.com/stretchr/testify/assert"
)

func TestWithdrawal(t *testing.T) {
	db := mocks_test.NewMockDB()
	db.Create(&database.Account{})
	a, err := account.GetByID(1, db)
	if !assert.Nil(t, err) {
		return
	}
	if !assert.NotNil(t, a) {
		return
	}

	balance.AddItems(a.ID, "minecraft:dirt", 64*4, db)
	for i := 0; i < 4; i++ {
		_, err := NewWithdrawal(WithdrawalOptions{
			AccountID: a.ID,
			Item:      "minecraft:dirt",
			Quantity:  64,
			Escrow:    "foo",
		}, db)
		assert.Nil(t, err)
	}
	var Withdrawals []database.Withdrawal
	err = db.Model(&database.Withdrawal{}).Where("account_id = ?", a.ID).Find(&Withdrawals).Error
	assert.Len(t, Withdrawals, 4)
	assert.Nil(t, err)

	b, err := balance.Get(a.ID, "minecraft:dirt", db)
	assert.Nil(t, err)
	assertBalance(t, b, 0)
}
