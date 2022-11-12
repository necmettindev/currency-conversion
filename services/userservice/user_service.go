package userservice

import (
	"errors"

	"github.com/necmettindev/currency-conversion/models/user"
	"github.com/necmettindev/currency-conversion/repositories/userrepo"

	"golang.org/x/crypto/bcrypt"
)

// UserService interface
type UserService interface {
	GetByUsername(username string) (*user.User, error)
	Create(*user.User) error
	HashPassword(rawPassword string) (string, error)
	ComparePassword(rawPassword string, passwordFromDB string) error
}

type userService struct {
	Repo   userrepo.UserRepo
	pepper string
}

func NewUserService(
	repo userrepo.UserRepo,
	pepper string) UserService {

	return &userService{
		Repo:   repo,
		pepper: pepper,
	}
}

func (us *userService) GetByUsername(username string) (*user.User, error) {
	if username == "" {
		return nil, errors.New("username(string) is required")
	}
	user, err := us.Repo.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userService) Create(user *user.User) error {
	hashedPass, err := us.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPass
	return us.Repo.Create(user)
}

func (us *userService) HashPassword(rawPassword string) (string, error) {
	passAndPepper := rawPassword + us.pepper
	hashed, err := bcrypt.GenerateFromPassword([]byte(passAndPepper), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), err
}

func (us *userService) ComparePassword(rawPassword string, passwordFromDB string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(passwordFromDB),
		[]byte(rawPassword+us.pepper),
	)
}
