package mocks_test

import (
	"time"

	"github.com/gin-gonic/gin"
)

func MockRouter(setup func(r *gin.Engine)) {
	r := gin.Default()
	setup(r)
	go func() {
		r.Run("127.0.0.1:8081")
	}()
	time.Sleep(time.Second)
}
