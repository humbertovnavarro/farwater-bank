package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/humbertovnavarro/farwater-bank/pkg/database"
	"github.com/humbertovnavarro/farwater-bank/pkg/middleware"
	mocks_test "github.com/humbertovnavarro/farwater-bank/pkg/mocks"
	"github.com/humbertovnavarro/farwater-bank/pkg/token"
	"github.com/stretchr/testify/assert"
)

func TestDeposit(t *testing.T) {
	db := mocks_test.NewMockDB()
	mocks_test.MockRouter(func(r *gin.Engine) {
		r.Use(func(ctx *gin.Context) {
			ctx.Set("db", db)
		})
		r.Use(middleware.AdminAuthentication)
		r.POST("/atm/deposit", Deposit)
	})
	db.Create(&database.Account{})
	body, err := json.Marshal(gin.H{
		"account_id": 1,
		"item":       "minecraft:dirt",
		"quantity":   32,
	})
	assert.Nil(t, err)
	req, err := http.NewRequest("POST", mocks_test.Route("atm", "deposit"), bytes.NewReader(body))
	assert.Nil(t, err)
	req.Header.Set("Content-Type", "application/json")
	tokenString, err := token.SignedString(token.AdminToken, "foo")
	assert.Nil(t, err)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tokenString))
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	respBytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	deposit := &database.Deposit{}
	json.Unmarshal(respBytes, deposit)
	assert.NotNil(t, deposit)
	assert.Equal(t, uint(1), deposit.ID)
	assert.Equal(t, "minecraft:dirt", deposit.Item)
	assert.Equal(t, uint64(32), deposit.Quantity)
}
