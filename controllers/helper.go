package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/necmettindev/currency-conversion/models/account"
	"github.com/necmettindev/currency-conversion/models/user"
)

// Response object as HTTP response
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// HTTPRes normalize HTTP Response format
func HTTPRes(c *gin.Context, httpCode int, msg string, data interface{}) {
	c.JSON(httpCode, Response{
		Code: httpCode,
		Msg:  msg,
		Data: data,
	})
}

func MapToAccountOutput(u *[]account.Account) *[]account.AccountOutput {

	accounts := make([]account.AccountOutput, len(*u))

	for i, acc := range *u {
		accounts[i] = account.AccountOutput{
			Currency: acc.Currency,
			Balance:  acc.Balance,
		}
	}

	return &accounts
}

func InputToUser(input user.UserRegisterInput) user.User {
	return user.User{
		Username:  input.Username,
		Password:  input.Password,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}
}

func MapToUserOutput(u *user.User) *user.UserOutput {
	return &user.UserOutput{
		ID:        u.ID,
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}
