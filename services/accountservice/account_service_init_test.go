package accountservice

import (
	"github.com/go-redis/redis/v9"
	"github.com/necmettindev/currency-conversion/configs"
	"github.com/necmettindev/currency-conversion/models/account"
	"github.com/stretchr/testify/mock"
)

var (
	testId = 1
)

type repoMock struct {
	mock.Mock
}

func (repo *repoMock) GetBalanceByID(id uint) (*[]account.Account, error) {
	args := repo.Called(id)
	return args.Get(0).(*[]account.Account), args.Error(1)
}

func (repo *repoMock) PostConversionCurrency(userId uint, firstCurrency string, secondCurrency string, amount float64, rate float64) (bool, error) {
	args := repo.Called(userId, firstCurrency, secondCurrency, amount, rate)
	return args.Bool(0), args.Error(1)
}

func (repo *repoMock) mockRedis() *redis.Client {
	redisConfig := configs.GetRedisConfig()
	s := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})

	return s
}
