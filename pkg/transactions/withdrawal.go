package transactions

import "gorm.io/gorm"

type Withdrawal struct {
	gorm.Model
	Item      string
	Amount    int64
	AccountID uint
	Escrow    string
	Signature string
}
