package mocks_test

import (
	"fmt"
	"os"

	"github.com/humbertovnavarro/farwater-bank/pkg/database"
)

const LISTEN = "127.0.0.1:8081"

// sets up a mock environment to test rest api responsses
func MockSetup() {
	os.Setenv("USER_SECRET", "user")
	os.Setenv("ADMIN_SECRET", "admin")
	os.Setenv("PEPPERS", "1,2,3,4")
	os.Setenv("LISTEN", LISTEN)
	database.New = NewMockDB
}

func Route(paths ...string) string {
	url := fmt.Sprintf("http://%s", LISTEN)
	if len(paths) == 0 {
		return url
	}
	for _, path := range paths {
		url = fmt.Sprintf("%s/%s", url, path)
	}
	return url
}
