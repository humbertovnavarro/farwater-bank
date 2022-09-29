package balance

import (
	"testing"

	"github.com/humbertovnavarro/farwater-bank/pkg/mocks"
)

func TestMain(m *testing.M) {
	mocks.MockSetup()
	m.Run()
}