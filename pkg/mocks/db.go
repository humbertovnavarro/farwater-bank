package mocks_test

import (
	"github.com/humbertovnavarro/farwater-bank/pkg/database"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewMockDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		logrus.Panic(err)
	}
	database.ApplySchema(db)
	return db
}
