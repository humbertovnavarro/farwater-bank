package transactions

import "gorm.io/gorm"

type Deposit struct {
	gorm.Model
	Item      string
	Amount    int64
	AccountID uint

	Escrow    string
	Signature string
}
