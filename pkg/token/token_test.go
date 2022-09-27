package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignedString(t *testing.T) {
	userSecret = []byte("user")
	adminSecret = []byte("admin")
	lastTokenType := AdminToken
	for i := 0; i < lastTokenType; i++ {
		tokenType := i
		tokenString, err := SignedString(tokenType, "bar")
		if err != nil {
			t.Fail()
		}
		assert.True(t, len(tokenString) > 3)
		token, err := ParseToken(tokenString, tokenType)
		assert.Nil(t, err)
		assert.Equal(t, token.Subject, "bar")
		assert.Equal(t, token.Type, tokenType)
	}
}
