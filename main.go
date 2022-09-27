package main

import (
	"os"

	"github.com/humbertovnavarro/farwater-bank/pkg/router"
)

func main() {
	router := router.New()
	router.Run(os.Getenv("LISTEN"))
}
