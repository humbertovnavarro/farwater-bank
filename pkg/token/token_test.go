package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignedString(t *testing.T) {
	userSecret = []byte("user")
	ATMSecret = []byte("atm")
	tokenString, err := SignedString(ATMToken, "bar")
	assert.Nil(t, err)
	token, err := ParseToken(tokenString)
	assert.Nil(t, err)
	assert.Equal(t, "bar", token.Subject)
	assert.Equal(t, 1, token.Type)

	tokenString, err = SignedString(UserToken, "bar")
	assert.Nil(t, err)
	token, err = ParseToken(tokenString)
	assert.Nil(t, err)
	assert.Equal(t, "bar", token.Subject)
	assert.Equal(t, 0, token.Type)
}
