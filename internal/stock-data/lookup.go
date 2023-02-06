package stocks

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/hmetcalfe/go-alpha-vantage/internal/urls"
)

func LookupStockInfo(symbol, apikey string, stock *Stock) error {
	logger := log.WithField("function", "stocks.ParseStockOverview()")

	url, err := urls.CompanyOverviewUrl(symbol, apikey)
	if err != nil {
		logText := fmt.Sprintf("failed to get the company information for the symbol %s", symbol)
		logger.WithError(err).Error(logText)

		return fmt.Errorf("%s: %w", logText, err)
	}

	resp, err := http.Get(url)
	if err != nil {
		logText := "failed to get information from the provided url"
		logger.WithError(err).Error(logText)

		return fmt.Errorf("%s: %w", logText, err)
	}

	defer resp.Body.Close()

	return nil
}
