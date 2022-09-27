package mocks

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	MinecraftUUID string
	Password      string
	PasswordSalt  string
	Pin           string
	PinSalt       string
	Frozen        bool
}

type Deposit struct {
	gorm.Model
	Item      string
	Amount    uint64
	AccountID uint
	Escrow    string
}

type Transfer struct {
	gorm.Model
	Item        string
	Amount      uint64
	AccountID   uint
	ToAccountID uint
	Escrow      string
}

type Withdrawal struct {
	gorm.Model
	Item      string
	Amount    uint64
	AccountID uint
	Escrow    string
}
