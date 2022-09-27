package itemcard

import (
	"testing"

	"github.com/humbertovnavarro/farwater-bank/pkg/account"
	"github.com/humbertovnavarro/farwater-bank/pkg/mocks"
	"github.com/humbertovnavarro/farwater-bank/pkg/token"
	"github.com/stretchr/testify/assert"
)

var db = mocks.NewMockDB()

func TestMain(m *testing.M) {
	db.AutoMigrate(account.Account{})
	db.AutoMigrate(ItemCard{})
	account.Register("notch", "1234", db)
	peppers = []string{"1", "2", "3", "4"}
	m.Run()
}

func Test(t *testing.T) {
	card, err := Issue(1, "1234", db)
	assert.Nil(t, err)
	assert.Equal(t, uint(1), card.AccountID)
	assert.NotEmpty(t, card.Salt)
	assert.NotEmpty(t, card.Token)
	token, err := token.ParseToken(card.Token, token.ItemCardToken)
	assert.Nil(t, err)
	assert.Equal(t, token.Subject, "1")
	assert.Nil(t, ValidatePin(card.ID, "1234", db))
}
