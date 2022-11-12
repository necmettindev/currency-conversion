package accountrepo

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/necmettindev/currency-conversion/models/account"
)

type AccountRepo interface {
	PostConversionCurrency(userId uint, firstCurrency string, secondCurrency string, amount float64, rate float64) (bool, error)
	GetBalanceByID(id uint) (*[]account.Account, error)
}

type accountRepo struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) AccountRepo {
	return &accountRepo{
		db: db,
	}
}

func (u *accountRepo) PostConversionCurrency(userId uint, firstCurrency string, secondCurrency string, amount float64, rate float64) (bool, error) {
	tx := u.db.Begin()
	acc := account.Account{}
	res := tx.Raw("SELECT * FROM accounts WHERE user_id = ? AND currency = ?", userId, firstCurrency).Scan(&acc)
	if res.Error != nil {
		tx.Rollback()
		return false, res.Error
	}
	if acc.Balance < amount {
		tx.Rollback()
		return false, errors.New("not enough balance")
	}
	updateFirstCurrencyBalance := tx.Exec("UPDATE accounts SET balance = balance - ? WHERE user_id = ? AND currency = ?", amount, userId, firstCurrency)
	if updateFirstCurrencyBalance.Error != nil {
		tx.Rollback()
		return false, res.Error
	}
	willBeUpdatedBalance := amount * rate
	updateSecondCurrencyBalance := tx.Exec("INSERT INTO accounts(currency, balance, user_id) VALUES(?, ?, ?) ON CONFLICT ON CONSTRAINT unique_currency_user_id DO UPDATE SET balance = accounts.balance + ?", secondCurrency, willBeUpdatedBalance, userId, willBeUpdatedBalance)
	if updateSecondCurrencyBalance.Error != nil {
		tx.Rollback()
		return false, res.Error
	}
	tx.Commit()
	u.db.Raw("INSERT INTO transactions (user_id, first_currency, second_currency, amount, rate) VALUES (?, ?, ?, ?, ?)", userId, firstCurrency, secondCurrency, amount, rate)
	return true, nil
}

func (u *accountRepo) GetBalanceByID(id uint) (*[]account.Account, error) {
	var accounts []account.Account
	resp := u.db.Raw("SELECT * FROM accounts WHERE user_id = ?", id).Scan(&accounts)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return &accounts, nil
}
