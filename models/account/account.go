package account

import "github.com/jinzhu/gorm"

type Account struct {
	gorm.Model
	ID       uint    `gorm:"primaryKey"`
	Currency string  `gorm:"size:255"`
	Balance  float64 `gorm:"size:255"`
	UserId   uint    `gorm:"NOT NULL"`
}

type PostCurrencyConversionInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AccountOutput struct {
	Currency string  `json:"currency"`
	Balance  float64 `json:"balance"`
}
