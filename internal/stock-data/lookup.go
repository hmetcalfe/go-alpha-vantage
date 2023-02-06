package stocks

import (
	"fmt"
	"io"
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

	// Perform the HTTP Get
	resp, err := http.Get(url)
	if err != nil {
		logText := "failed to get information from the provided url"
		logger.WithError(err).Error(logText)

		return fmt.Errorf("%s: %w", logText, err)
	}

	// Defer the response body close until function completion
	defer resp.Body.Close()

	// Read in the response data
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logText := "failed to read in the response body"
		logger.WithError(err).Error(logText)

		return fmt.Errorf("%s: %w", logText, err)
	}

	// Parse the json into the Stock structure
	*stock, err = ParseStockOverview(body)
	if err != nil {
		logText := "failed to parse the stock json"
		logger.WithError(err).Error(logText)

		return fmt.Errorf("%s: %w", logText, err)
	}

	return nil
}
