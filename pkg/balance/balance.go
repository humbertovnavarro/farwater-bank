package balance

import "gorm.io/gorm"

type Balance struct {
	gorm.Model
	Item     string
	Quantity uint64
}
