package main

import (
	"github.com/humbertovnavarro/farwater-bank/pkg/router"
)

func main() {
	router := router.New()
	router.Run("127.0.0.1:8080")
}
