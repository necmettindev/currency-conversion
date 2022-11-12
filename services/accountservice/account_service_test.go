package accountservice

import (
	"testing"

	"github.com/necmettindev/currency-conversion/models/account"

	"github.com/stretchr/testify/assert"
)

func TestGetBalanceById(t *testing.T) {
	t.Run("GetBalanceById", func(t *testing.T) {
		var expected *[]account.Account = &[]account.Account{
			{
				ID:       1,
				Currency: "USD",
				Balance:  500.000000,
				UserId:   1,
			},
			{
				ID:       2,
				Currency: "TRY",
				Balance:  100.000000,
				UserId:   1,
			},
		}
		accountRepo := new(repoMock)
		redis := accountRepo.mockRedis()
		accountRepo.On("GetBalanceByID", uint(testId)).Return(expected, nil)
		accountService := NewAccountService(redis, accountRepo)
		result, err := accountService.GetBalanceByID(uint(testId))
		if err != nil {
			t.Errorf("GetBalanceByID returned an error: %s", err)
		}

		if (*result)[0].Balance != (*expected)[0].Balance {
			t.Errorf("GetBalanceByID returned wrong result: %f", (*result)[0].Balance)
		}

		assert.Equal(t, (*result)[0].Currency, (*expected)[0].Currency)
		assert.Nil(t, err)
	})

}
