package user

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	FirstName string `gorm:"size:255"`
	LastName  string `gorm:"size:255"`
	Username  string `gorm:"NOT NULL; UNIQUE_INDEX"`
	Password  string `gorm:"NOT NULL"`
}

type UserRegisterInput struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
}

type UserLoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserOutput struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
}
