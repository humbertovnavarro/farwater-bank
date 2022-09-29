package balance

import (
	"testing"

	"github.com/humbertovnavarro/farwater-bank/pkg/database"
	"github.com/humbertovnavarro/farwater-bank/pkg/mocks"
	"github.com/stretchr/testify/assert"
)

func assertBalance(t *testing.T, b *Balance, quantity uint64) {
	if !assert.NotNil(t, b) {
		return
	}
	assert.Equal(t, "minecraft:dirt", b.Item)
	assert.Equal(t, quantity, b.Quantity)
}

func TestNew(t *testing.T) {
	db := mocks.NewMockDB()
	db.Create(&database.Account{})
	b, err := New(1, "minecraft:dirt", 64, db)
	assert.Nil(t, err)
	assertBalance(t, b, 64)
	// pluck straight from db
	databaseBalance := &database.Balance{}
	assert.Nil(t, db.Model(&database.Balance{}).First(databaseBalance).Error)
	assertBalance(t, &Balance{
		*databaseBalance,
	}, 64)
}

func TestGet(t *testing.T) {
	db := mocks.NewMockDB()
	db.Create(&database.Account{})
	New(1, "minecraft:dirt", 64, db)
	b, err := Get(1, "minecraft:dirt", db)
	assert.Nil(t, err)
	assert.NotNil(t, b)
	assert.Equal(t, "minecraft:dirt", b.Item)
	assert.Equal(t, uint64(64), b.Quantity)
}

func TestAddItems(t *testing.T) {
	db := mocks.NewMockDB()
	db.Create(&database.Account{})
	err := AddItems(1, "minecraft:dirt", 64, db)
	assert.Nil(t, err)
	b, err := Get(1, "minecraft:dirt", db)
	assert.Nil(t, err)
	assertBalance(t, b, 64)

	err = AddItems(1, "minecraft:dirt", 64, db)
	assert.Nil(t, err)
	b, err = Get(1, "minecraft:dirt", db)
	assert.Nil(t, err)
	assertBalance(t, b, 128)
}

func TestRemoveItems(t *testing.T) {
	db := mocks.NewMockDB()
	db.Create(&database.Account{})

	err := RemoveItems(1, "minecraft:dirt", 64, db)
	assert.Error(t, err)
	assert.Equal(t, "record not found", err.Error())

	err = AddItems(1, "minecraft:dirt", 66, db)
	assert.Nil(t, err)

	err = RemoveItems(1, "minecraft:dirt", 2, db)
	assert.Nil(t, err)
	b, err := Get(1, "minecraft:dirt", db)
	assert.Nil(t, err)
	assertBalance(t, b, 64)

	err = RemoveItems(1, "minecraft:dirt", 99999, db)
	assert.NotNil(t, err)
	assert.Equal(t, "1 attempted to overdraft minecraft:dirt", err.Error())

}
