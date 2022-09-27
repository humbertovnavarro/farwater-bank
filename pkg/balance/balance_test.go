package balance

import (
	"testing"

	"github.com/humbertovnavarro/farwater-bank/pkg/mocks"
	"gorm.io/gorm"
)

var db *gorm.DB = mocks.NewMockDB()

func TestMain(m *testing.M) {
	m.Run()
}
