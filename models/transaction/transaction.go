package transaction

import "github.com/jinzhu/gorm"

type Transaction struct {
	gorm.Model
	ID             uint    `gorm:"primaryKey"`
	UserId         uint    `gorm:"NOT NULL"`
	FirstCurrency  string  `gorm:"size:255"`
	SecondCurrency string  `gorm:"size:255"`
	Amount         float64 `gorm:"size:255"`
	Rate           float64 `gorm:"size:255"`
}
