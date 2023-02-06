package stocks

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	data "github.com/hmetcalfe/go-alpha-vantage/internal/stock-data"
)

func GetStockInfo(symbol, apikey string) (data.Stock, error) {
	logger := log.WithField("function", "stocks.GetStockInfo()")
	var stock data.Stock

	err := data.LookupStockInfo(symbol, apikey, &stock)
	if err != nil {
		logText := fmt.Sprintf("failed to lookup stock information %s", symbol)
		logger.WithError(err).Error(logText)

		return data.Stock{}, fmt.Errorf("%s: %w", logText, err)
	}

	return stock, nil
}
