package financeservice

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	financeapi "github.com/necmettindev/currency-conversion/models/finance"
)

type FinanceService interface {
	GetCurrenciesMarketPrice(firstCurrency string, secondCurrency string) (*financeapi.FinanceAPI, error)
}

type financeService struct {
	http *http.Client
}

func (u *financeService) GetCurrenciesMarketPrice(firstCurrency string, secondCurrency string) (*financeapi.FinanceAPI, error) {
	u.http.Timeout = time.Second * 10
	resp, err := u.http.Get("https://query1.finance.yahoo.com/v7/finance/spark?symbols=" + strings.ToUpper(firstCurrency) + strings.ToUpper(secondCurrency) + "%3DX")

	if err != nil {
		return nil, err
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("error")
	}

	body, readErr := io.ReadAll(resp.Body)

	if readErr != nil {
		log.Fatal(readErr)
	}

	result := financeapi.FinanceAPI{}
	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return &result, nil

}

func NewFinanceService() FinanceService {
	return &financeService{
		http: &http.Client{},
	}
}
