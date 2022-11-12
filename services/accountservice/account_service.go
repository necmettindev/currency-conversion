package accountservice

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/necmettindev/currency-conversion/configs"
	"github.com/necmettindev/currency-conversion/models/account"
	"github.com/necmettindev/currency-conversion/repositories/accountrepo"
	"github.com/necmettindev/currency-conversion/services/financeservice"
)

type AccountService interface {
	GetExchangeRate(firstCurrency string, secondCurrency string, financeService financeservice.FinanceService) (float64, error)
	CheckExchangeRate(firstCurrency string, secondCurrency string) (float64, error)
	PostConversionCurrency(userId uint, firstCurrency string, secondCurrency string, amount float64) (bool, error)
	GetBalanceByID(id uint) (*[]account.Account, error)
}

type accountService struct {
	rdb  *redis.Client
	repo accountrepo.AccountRepo
}

var ctx = context.Background()

func NewAccountService(
	rdb *redis.Client,
	Repo accountrepo.AccountRepo,
) AccountService {
	return &accountService{
		rdb:  rdb,
		repo: Repo,
	}
}

func (account *accountService) GetExchangeRate(firstCurrency string, secondCurrency string, financeService financeservice.FinanceService) (float64, error) {
	val, err := account.rdb.Get(ctx, firstCurrency+secondCurrency).Result()

	config := configs.GetConfig()

	if err == redis.Nil {
		result, err := financeService.GetCurrenciesMarketPrice(firstCurrency, secondCurrency)
		marketPrice := result.Spark.Result[0].Response[0].Meta.RegularMarketPrice
		marketPriceWithFee := marketPrice - (marketPrice * config.FeePercentage)
		if err != nil {
			return 0, err
		}
		account.rdb.Set(ctx, firstCurrency+secondCurrency, marketPriceWithFee, time.Minute*3)
		return marketPriceWithFee, nil
	} else if err != nil {
		return 0, err
	} else {
		parsedVal, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return 0, err
		}
		return parsedVal, err
	}
}

func (account *accountService) CheckExchangeRate(firstCurrency string, secondCurrency string) (float64, error) {
	val, err := account.rdb.Get(ctx, firstCurrency+secondCurrency).Result()

	if err == redis.Nil {
		return 0, errors.New("no data")
	} else if err != nil {
		return 0, err
	} else {
		parsedVal, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return 0, err
		}
		return parsedVal, err
	}
}

func (account *accountService) PostConversionCurrency(userId uint, firstCurrency string, secondCurrency string, amount float64) (bool, error) {
	val, err := account.rdb.Get(ctx, firstCurrency+secondCurrency).Result()

	if err == redis.Nil {
		return false, errors.New("no data")
	} else if err != nil {
		return false, errors.New("no data")
	}

	parsedVal, err := strconv.ParseFloat(val, 64)

	if err != nil {
		return false, err
	}

	result, err := account.repo.PostConversionCurrency(userId, firstCurrency, secondCurrency, amount, parsedVal)

	if err != nil {
		return false, err
	}

	return result, nil
}

func (account *accountService) GetBalanceByID(id uint) (*[]account.Account, error) {
	if id == 0 {
		return nil, errors.New("id param is required")
	}

	result, err := account.repo.GetBalanceByID(id)

	if err != nil {
		return nil, err
	}

	return result, nil
}
