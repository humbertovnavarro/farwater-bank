package transactions

import "gorm.io/gorm"

type TransactionSignature struct {
	gorm.Model
	TransactionID uint
	Signature     string
}
