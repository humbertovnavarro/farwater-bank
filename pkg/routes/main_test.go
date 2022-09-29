package routes

import (
	"testing"

	mocks_test "github.com/humbertovnavarro/farwater-bank/pkg/mocks"
)

func TestMain(m *testing.M) {
	mocks_test.MockSetup()
	m.Run()
}
