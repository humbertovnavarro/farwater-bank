package transactions

import "gorm.io/gorm"

type Transfer struct {
	gorm.Model
	Item        string
	Amount      int64
	AccountID   uint
	ToAccountID uint

	Escrow    string
	Signature string
}
