package urls

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/hmetcalfe/go-alpha-vantage/internal/utils"
)

const (
	ALPHA_VANTAGE_URL = "https://www.alphavantage.co/query?"
)

const (
	OVERVIEW_FUNCTION = "OVERVIEW"
)

func APIUrl() string {
	return ALPHA_VANTAGE_URL
}

func CompanyOverviewUrl(symbol, apikey string) (string, error) {
	logger := log.WithField("function", "urls.CompanyOverviewUrl()")

	// As a security precaution, the API Key is passed through functions as a b64 string
	key, err := utils.DecodeAPIKey(apikey)
	if err != nil {
		logText := fmt.Sprintf("failed to decode key: %s while fetching %s overview", key, symbol)
		logger.WithError(err).Error(logText)

		return "", fmt.Errorf("%s: %w", logText, err)
	}

	// Format the API URL https://www.alphavantage.co/query?function=FUNCTION&symbol=SYMBOL&apikey=apikey
	return fmt.Sprintf("%sfunction=%s&symbol=%s&apikey=%s", ALPHA_VANTAGE_URL, OVERVIEW_FUNCTION, symbol, key), nil
}
