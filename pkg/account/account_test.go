package account

import (
	"testing"

	"github.com/humbertovnavarro/farwater-bank/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var notchUUID = "069a79f444e94726a5befca90e38aaf5"
var db *gorm.DB = mocks.NewMockDB()

func TestMain(m *testing.M) {
	m.Run()
}

func TestRegister(t *testing.T) {
	// register
	a, err := Register("notch", "1234", "1234", db)
	assert.Nil(t, err)
	assert.Equal(t, uint(1), a.ID)
	assert.Equal(t, a.MinecraftUUID, notchUUID)
	// verify its there
	a, err = GetByUUID(a.MinecraftUUID, db)
	assert.Nil(t, err)
	assert.Equal(t, uint(1), a.ID)
	assert.Equal(t, notchUUID, a.MinecraftUUID)
	// test passwords

	// try to register account that doesn't exist
	a, err = Register("_", "1234", "1234", db)
	assert.Error(t, err)
	assert.Equal(t, "could not fetch minecraft uuid", err.Error())
	assert.Nil(t, a)
}
