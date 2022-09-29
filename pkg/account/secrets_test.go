package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashSecret(t *testing.T) {
	salt, hex, err := HashSecret("I like pineapples on pizza")
	assert.Nil(t, err)
	assert.Len(t, salt, 36)
	err = VerifySecret("I like pineapples on pizza", hex, salt)
	assert.Nil(t, err)
}
