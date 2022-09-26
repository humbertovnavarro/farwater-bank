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
	m.Run()
}
func TestIssue(t *testing.T) {
	a, err := account.Register("notch", "1234", db)
	assert.Nil(t, err)
	assert.NotNil(t, a)
	card, err := Issue(a.ID, "salt", db)
	assert.Nil(t, err)
	assert.NotNil(t, card)
	assert.Equal(t, a.ID, card.AccountID)
	assert.Equal(t, card.Frozen, false)
	tokenString, _ := token.SignedString(token.ItemCardToken, "1")
	assert.Equal(t, card.Token, tokenString)
}
