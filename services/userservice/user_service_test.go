package userservice

import (
	"errors"
	"fmt"
	"testing"

	"github.com/necmettindev/currency-conversion/models/user"

	"github.com/stretchr/testify/assert"
)

func TestGetByUsername(t *testing.T) {
	t.Run("Get a user", func(t *testing.T) {
		expected := &user.User{
			FirstName: "Test",
			LastName:  "User",
		}

		userRepo := new(repoMock)
		u := NewUserService(userRepo, pepper)
		userRepo.On("GetByUsername", testUsername).Return(expected, nil)

		result, _ := u.GetByUsername(testUsername)

		assert.EqualValues(t, expected, result)
	})

	t.Run("Get error if username is empty", func(t *testing.T) {
		expected := errors.New("username(string) is required")

		userRepo := new(repoMock)
		u := NewUserService(userRepo, pepper)

		result, err := u.GetByUsername("")

		assert.Nil(t, result)
		assert.EqualValues(t, expected, err)
	})

	t.Run("Get error if it has error", func(t *testing.T) {
		expected := errors.New("Nop")

		userRepo := new(repoMock)
		u := NewUserService(userRepo, pepper)
		userRepo.On("GetByUsername", testUsername).Return(&user.User{}, expected)

		result, err := u.GetByUsername(testUsername)

		assert.Nil(t, result)
		assert.EqualValues(t, expected, err)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create a user", func(t *testing.T) {
		usr := &user.User{
			Username: "necmettin",
			Password: "abc123",
		}

		userRepo := new(repoMock)
		u := NewUserService(userRepo, pepper)
		userRepo.On("Create", usr).Return(nil)

		result := u.Create(usr)

		assert.Nil(t, result)
	})

	t.Run("Create a user fails", func(t *testing.T) {
		err := errors.New(("oops"))
		usr := &user.User{
			Username: "necmettin",
		}

		userRepo := new(repoMock)
		u := NewUserService(userRepo, pepper)

		userRepo.On("Create", usr).Return(err)
		result := u.Create(usr)

		assert.EqualValues(t, result, err)
	})
}

func TestComparePassword(t *testing.T) {
	t.Run("match password", func(t *testing.T) {
		testPass := "test123"

		userRepo := new(repoMock)
		u := NewUserService(userRepo, pepper)

		hashedPass, err := u.HashPassword(testPass)
		fmt.Println(hashedPass, err)
		err = u.ComparePassword(testPass, hashedPass)
		assert.Nil(t, err)
	})

	t.Run("not match password", func(t *testing.T) {
		testPass := "test123"

		userRepo := new(repoMock)
		u := NewUserService(userRepo, pepper)

		hashedPass, err := u.HashPassword(testPass)
		fmt.Println(hashedPass, err)
		err = u.ComparePassword("test1234", hashedPass)
		assert.NotNil(t, err)
	})
}
