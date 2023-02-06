package stocks

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hmetcalfe/go-alpha-vantage/internal/utils"
)

func TestLookupStockInfo(t *testing.T) {
	apikey := "demo"
	symbol := "IBM"

	// Encode the apikey
	apikey = utils.EncodeAPIKey(apikey)

	// Lookup IBM
	var stock Stock
	err := LookupStockInfo(symbol, apikey, &stock)
	assert.NoError(t, err)

	symbol = "asdasdasd"
	apikey = "asdasd"
	err = LookupStockInfo(symbol, apikey, &stock)
	assert.Error(t, err)
}
