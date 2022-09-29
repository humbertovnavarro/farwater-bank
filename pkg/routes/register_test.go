package routes

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/humbertovnavarro/farwater-bank/pkg/account"
	"github.com/humbertovnavarro/farwater-bank/pkg/middleware"
	mocks_test "github.com/humbertovnavarro/farwater-bank/pkg/mocks"
	"github.com/humbertovnavarro/farwater-bank/pkg/token"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	db := mocks_test.NewMockDB()
	mocks_test.MockRouter(func(r *gin.Engine) {
		r.Use(func(ctx *gin.Context) {
			ctx.Set("db", db)
		})
		r.Use(middleware.AdminAuthentication)
		r.POST("/atm/register", Register)
	})
	body, err := json.Marshal(gin.H{
		"username": "Notch",
		"password": "test",
		"pin":      "1234",
	})
	assert.Nil(t, err)

	// No auth
	resp, err := http.Post(mocks_test.Route("atm", "register"), "application/json", bytes.NewReader(body))
	if !assert.Nil(t, err) {
		fmt.Println(err)
		return
	}
	expected, err := json.Marshal(gin.H{
		"error": "unauthorized",
	})
	assert.Nil(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	body, err = io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, string(expected), string(body))

	// With Auth
	body, err = json.Marshal(gin.H{
		"username": "Notch",
		"password": "test",
		"pin":      "1234",
	})
	assert.Nil(t, err)
	req, err := http.NewRequest("POST", mocks_test.Route("atm", "register"), bytes.NewReader(body))
	assert.Nil(t, err)
	token, err := token.SignedString(token.AdminToken, "foo")
	assert.Nil(t, err)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	client := http.DefaultClient
	resp, err = client.Do(req)
	assert.Nil(t, err)
	body, err = io.ReadAll(resp.Body)
	assert.Nil(t, err)
	expected, err = json.Marshal(gin.H{
		"account_id":     1,
		"minecraft_uuid": "069a79f444e94726a5befca90e38aaf5",
	})
	assert.Nil(t, err)
	assert.Equal(t, string(expected), string(body))

	a, err := account.GetByID(1, db)
	assert.Nil(t, err)
	assert.Equal(t, "069a79f444e94726a5befca90e38aaf5", a.MinecraftUUID)
}
