package database

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
	Quantity  uint64
	AccountID uint
	Escrow    string
}

type Transfer struct {
	gorm.Model
	Item        string
	Quantity    uint64
	AccountID   uint
	ToAccountID uint
	Escrow      string
}

type Withdrawal struct {
	gorm.Model
	Item      string
	Quantity  uint64
	AccountID uint
	Escrow    string
}

type Balance struct {
	gorm.Model
	AccountID uint
	Item      string
	Quantity  uint64
}

func ApplySchema(db *gorm.DB) {
	db.AutoMigrate(Account{})
	db.AutoMigrate(Deposit{})
	db.AutoMigrate(Transfer{})
	db.AutoMigrate(Withdrawal{})
	db.AutoMigrate(Balance{})
}
