package routes

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/humbertovnavarro/farwater-bank/pkg/account"
	"github.com/humbertovnavarro/farwater-bank/pkg/token"
	"gorm.io/gorm"
)

type VerifyPinRequest struct {
	Pin          string `json:"pin"`
	Identity     string `json:"identity"`
	IdentityType string `json:"identity_type"`
}

func (r *VerifyPinRequest) Valid() error {
	if len(r.Pin) < 4 {
		return errors.New("invalid pin")
	}
	if !(r.IdentityType == "uuid" || r.IdentityType == "id") {
		return errors.New("invalid identity type")
	}
	return nil
}

func VerifyPin(c *gin.Context) {
	authorization := c.MustGet("authorization").(*token.Token)
	if !(authorization.Type == token.AdminToken) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}
	request := &VerifyPinRequest{}
	if err := c.BindJSON(request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
	}
	err := request.Valid()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	db := c.MustGet("db").(*gorm.DB)
	switch request.IdentityType {
	case "uuid":
		a, err := account.GetByUUID(request.Identity, db)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}
		if err := a.VerifyPin(request.Pin); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}
	case "id":
		id, err := strconv.ParseUint(request.Identity, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "id must be a positive integer",
			})
			return
		}
		a, err := account.GetByID(uint(id), db)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}
		if err := a.VerifyPin(request.Pin); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}
	}
	c.AbortWithStatus(http.StatusOK)
}
