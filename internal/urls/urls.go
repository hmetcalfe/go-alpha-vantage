package urls

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/hmetcalfe/go-alpha-vantage/internal/utils"
)

const (
	alpha_vantage_url = "https://www.alphavantage.co/query?"
)

const (
	overview_function = "OVERVIEW"
)

func APIUrl() string {
	return alpha_vantage_url
}

func OverviewFunction() string {
	return overview_function
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
	return fmt.Sprintf("%sfunction=%s&symbol=%s&apikey=%s", APIUrl(), OverviewFunction(), symbol, key), nil
}
