package minecraft

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchUsername(t *testing.T) {
	username, err := FetchUsername("069a79f444e94726a5befca90e38aaf5")
	assert.Nil(t, err)
	assert.Equal(t, "Notch", username)
}

func TestFetchUUID(t *testing.T) {
	username, err := FetchUUID("notch")
	assert.Nil(t, err)
	assert.Equal(t, "069a79f444e94726a5befca90e38aaf5", username)
}
