package userservice

import (
	"github.com/necmettindev/currency-conversion/models/user"
	"github.com/stretchr/testify/mock"
)

var (
	pepper       = "pepper"
	testUsername = "necmettin"
)

type repoMock struct {
	mock.Mock
}

func (repo *repoMock) GetByUsername(username string) (*user.User, error) {
	args := repo.Called(username)
	return args.Get(0).(*user.User), args.Error(1)
}

func (repo *repoMock) Create(user *user.User) error {
	args := repo.Called(user)
	return args.Error(0)
}
