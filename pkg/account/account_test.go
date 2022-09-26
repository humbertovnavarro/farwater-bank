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
	db.AutoMigrate(&Account{})
	m.Run()
}

func TestRegister(t *testing.T) {
	// register
	account, err := Register("notch", "1234", db)
	assert.Nil(t, err)
	assert.Equal(t, uint(1), account.ID)
	assert.Equal(t, account.MinecraftUUID, notchUUID)
	// verify its there
	account, err = GetByID(account.ID, db)
	assert.Nil(t, err)
	assert.Equal(t, uint(1), account.ID)
	assert.Equal(t, notchUUID, account.MinecraftUUID)

	// try to register account that doesn't exist
	account, err = Register("_", "1234", db)
	assert.Error(t, err)
	assert.Equal(t, "could not fetch minecraft uuid", err.Error())
	assert.Nil(t, account)
}

func TestFetchMinecraftUsername(t *testing.T) {
	username, err := FetchMinecraftUsername("069a79f444e94726a5befca90e38aaf5")
	assert.Nil(t, err)
	assert.Equal(t, "Notch", username)
}

func TestGetByUUID(t *testing.T) {
	Register("dinnerbone", "1234", db)
	uuid, _ := FetchMinecraftUUID("dinnerbone")
	account, err := GetByUUID(uuid, db)
	assert.Nil(t, err)
	assert.Equal(t, uuid, account.MinecraftUUID)
}
