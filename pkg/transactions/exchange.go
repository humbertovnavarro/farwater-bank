package transactions

import "gorm.io/gorm"

type Exchange struct {
	gorm.Model
	Item         string
	Amount       int64
	AccountID    uint
	ForItem      string
	ForAmount    int64
	ForAccountID uint

	Escrow    string
	Signature string
}
