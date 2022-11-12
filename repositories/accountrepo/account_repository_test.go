package accountrepo

import (
	"database/sql/driver"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/necmettindev/currency-conversion/models/account"
	"github.com/stretchr/testify/assert"
)

func setupDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("can't create sqlmock: %s", err)
	}

	gormDB, gerr := gorm.Open("postgres", db)
	if gerr != nil {
		log.Fatalf("can't open gorm connection: %s", err)
	}
	gormDB.LogMode(true)
	return gormDB, mock
}

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

// func TestGetConversionCurrency(t *testing.T) {
// 	gormDB, mock := setupDB()
// 	defer gormDB.Close()

// 	t.Run("GetConversionCurrency", func(t *testing.T) {
// 		rows := sqlmock.NewRows([]string{"id", "currency", "balance", "user_id"}).
// 			AddRow(1, "TRY", 100.0, 1).
// 			AddRow(2, "USD", 200.0, 1)

// 		mock.ExpectQuery("SELECT (.+) FROM accounts WHERE user_id = ?").
// 			WithArgs(1).
// 			WillReturnRows(rows)

// 		accRepo := NewAccountRepo(gormDB)
// 		result, err := accRepo.PostConversionCurrency(1, "TRY", "USD", 100.0, 0.2)
// 		if err != nil {
// 			t.Errorf("GetConversionCurrency returned an error: %s", err)
// 		}
// 		if result != true {
// 			t.Errorf("GetConversionCurrency returned wrong result: %t", result)
// 		}

// 	})

// }

func TestGetBalanceByID(t *testing.T) {
	gormDB, mock := setupDB()
	defer gormDB.Close()

	t.Run("GetBalanceByID", func(t *testing.T) {
		var expected = &[]account.Account{
			{
				ID:       1,
				Currency: "USD",
				Balance:  500.0,
				UserId:   1,
			},
		}

		u := NewAccountRepo(gormDB)
		sqlStr := "SELECT * FROM accounts WHERE user_id = $1"

		rows := sqlmock.NewRows([]string{"id", "currency", "balance", "user_id"}).
			AddRow(1, "USD", 500.000000, 1)

		mock.ExpectQuery(regexp.QuoteMeta(sqlStr)).
			WithArgs(1).
			WillReturnRows(rows)

		result, err := u.GetBalanceByID(1)
		if err != nil {
			t.Errorf("GetBalanceByID returned an error: %s", err)
		}

		if (*result)[0].Balance != (*expected)[0].Balance {
			t.Errorf("GetBalanceByID returned wrong result: %f", (*result)[0].Balance)
		}

		assert.Equal(t, (*result)[0].Currency, (*expected)[0].Currency)
		assert.Nil(t, mock.ExpectationsWereMet())

	})

	t.Run("GetBalanceByIDError", func(t *testing.T) {
		mock.ExpectQuery("SELECT (.+) FROM accounts WHERE user_id = ?").
			WithArgs(1).
			WillReturnError(sqlmock.ErrCancelled)

		accRepo := NewAccountRepo(gormDB)
		accounts, err := accRepo.GetBalanceByID(1)
		if err == nil {
			t.Errorf("GetBalanceByID should return an error")
		}
		if accounts != nil {
			t.Errorf("GetBalanceByID should return nil")
		}
	})

}
