package financeservice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrenciesMarketPrice(t *testing.T) {
	// Arrange
	financeService := NewFinanceService()
	// Act
	result, err := financeService.GetCurrenciesMarketPrice("usd", "try")
	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "USDTRY=X", result.Spark.Result[0].Symbol)

}
