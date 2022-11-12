package userrepo

import (
	"github.com/necmettindev/currency-conversion/models/user"

	"github.com/jinzhu/gorm"
)

type UserRepo interface {
	GetByUsername(username string) (*user.User, error)
	Create(user *user.User) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) GetByUsername(username string) (*user.User, error) {
	var user user.User
	if err := u.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) Create(user *user.User) error {
	return u.db.Create(user).Error
}
