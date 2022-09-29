package account

import (
	"testing"

	mocks_test "github.com/humbertovnavarro/farwater-bank/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var notchUUID = "069a79f444e94726a5befca90e38aaf5"
var db *gorm.DB = mocks_test.NewMockDB()

func TestRegister(t *testing.T) {
	// register
	_, err := Register("notch", "1234", "1234", db)
	assert.Nil(t, err)
	aA, err := GetByID(1, db)
	assert.Nil(t, err)
	assert.Equal(t, aA.MinecraftUUID, notchUUID)
	// verify its there
	aB, err := GetByUUID(aA.MinecraftUUID, db)
	assert.Nil(t, err)
	assert.Equal(t, aA.ID, aB.ID)
	assert.Equal(t, notchUUID, aB.MinecraftUUID)
	// test passwords

	// try to register account that doesn't exist
	_, err = Register("_", "1234", "1234", db)
	assert.Error(t, err)
	assert.Equal(t, "could not fetch minecraft uuid", err.Error())

}
