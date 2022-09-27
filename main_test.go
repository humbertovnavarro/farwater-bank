package main

import (
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/humbertovnavarro/farwater-bank/pkg/mocks"
	"github.com/stretchr/testify/assert"
)

const LISTEN = "127.0.0.1:8081"

func TestMain(m *testing.M) {
	mocks.MockSetup()
}

func TestSmoke(t *testing.T) {
	go func() {
		main()
	}()
	time.Sleep(1 * time.Second)
	resp, err := http.Get(mocks.Route())
	if !assert.Nil(t, err) {
		fmt.Println(err)
		return
	}
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "404 page not found", string(body))
}
