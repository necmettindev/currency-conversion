package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	accountservice "github.com/necmettindev/currency-conversion/services/accountservice"
	"github.com/necmettindev/currency-conversion/services/financeservice"
	"github.com/necmettindev/currency-conversion/services/userservice"
)

type AccountController interface {
	GetCurrencyConversionRate(*gin.Context)
	PostCurrencyConversion(*gin.Context)
	GetBalances(*gin.Context)
}

type accountController struct {
	us  userservice.UserService
	acs accountservice.AccountService
	fs  financeservice.FinanceService
}

func NewAccountController(us userservice.UserService, acs accountservice.AccountService, fs financeservice.FinanceService) AccountController {
	return &accountController{
		acs: acs,
		us:  us,
		fs:  fs,
	}
}

// @Security ApiKeyAuth
// @param Authorization header string true "Bearer {token}"
// @BasePath /v1
// GetCurrencyConversionRate godoc
// @Summary Get currency conversion rate
// @Description This endpoint is used to get currency conversion rate
// @Tags Account
// @Accept json
// @Produce json
// @Param first_currency path string true "First Currency"
// @Param second_currency path string true "Second Currency"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /accounts/{first_currency}/{second_currency}/rate [get]
func (ctl *accountController) GetCurrencyConversionRate(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		HTTPRes(c, http.StatusUnauthorized, "Invalid user id", gin.H{
			"error": "Invalid user id",
			"userI": userId,
		})
		return
	}

	firstCurrency := c.Param("first_currency")
	secondCurrency := c.Param("second_currency")
	marketPrice, err := ctl.acs.GetExchangeRate(firstCurrency, secondCurrency, ctl.fs)
	if err != nil {
		HTTPRes(c, http.StatusInternalServerError, "Error", nil)
		return
	}

	HTTPRes(c, http.StatusOK, "Success", gin.H{
		"success":      true,
		"error":        0,
		"market_price": marketPrice,
	})

}

// @Security ApiKeyAuth
// @param Authorization header string true "Bearer {token}"
// @BasePath /v1
// PostCurrencyConversion godoc
// @Summary Post Currency Conversion
// @Description Post Currency Conversion
// @Tags Account
// @Accept  json
// @Produce  json
// @Param first_currency path string true "First Currency"
// @Param second_currency path string true "Second Currency"
// @Param amount path string true "Amount"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /accounts/{first_currency}/{second_currency}/{amount}/conversion [post]
func (ctl *accountController) PostCurrencyConversion(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		HTTPRes(c, http.StatusUnauthorized, "Invalid user id", gin.H{
			"error": "Invalid user id",
			"userI": userId,
		})
		return
	}

	firstCurrency := c.Param("first_currency")
	secondCurrency := c.Param("second_currency")
	amount, err := strconv.ParseFloat(c.Param("amount"), 64)
	if err != nil {
		HTTPRes(c, http.StatusBadRequest, "Invalid amount", gin.H{
			"error": "Invalid amount",
		})
		return
	}
	marketPrice, err := ctl.acs.CheckExchangeRate(firstCurrency, secondCurrency)

	if err != nil {
		HTTPRes(c, http.StatusBadRequest, "Offer expired", gin.H{
			"error":        err.Error(),
			"market_price": marketPrice,
		})
		return
	}

	result, err := ctl.acs.PostConversionCurrency(userId.(uint), firstCurrency, secondCurrency, amount)
	if err != nil {
		HTTPRes(c, http.StatusInternalServerError, "Error", gin.H{
			"success": false,
			"error":   1,
			"message": err.Error(),
		})
		return
	}

	HTTPRes(c, http.StatusOK, "Success", gin.H{
		"success": result,
		"error":   0,
		"message": "Currency conversion is successful",
	})
}

// @Security ApiKeyAuth
// @param Authorization header string true "Bearer {token}"
// @BasePath /v1
// GetBalances godoc
// @Summary Get Balances
// @Description Get Balances
// @Tags Account
// @Accept  json
// @Produce  json
// @Success 200 {object} Response
// @Failure 500 {object} Response
// @Router /accounts [get]
func (ctl *accountController) GetBalances(c *gin.Context) {
	id, exists := c.Get("user_id")
	if !exists {
		HTTPRes(c, http.StatusBadRequest, "Invalid User ID", nil)
		return
	}

	resp, err := ctl.acs.GetBalanceByID(id.(uint))

	if err != nil {
		HTTPRes(c, http.StatusInternalServerError, "Error", nil)
		return
	}

	accounts := MapToAccountOutput(resp)

	HTTPRes(c, http.StatusOK, "Success", accounts)

}
