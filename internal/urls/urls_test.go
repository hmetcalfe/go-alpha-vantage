package urls

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hmetcalfe/go-alpha-vantage/internal/utils"
)

func TestCompanyOverviewUrl(t *testing.T) {
	apikey := "demo"
	symbol := "IBM"

	expectedUrl := fmt.Sprintf("%sfunction=%s&symbol=%s&apikey=%s", APIUrl(), OverviewFunction(), symbol, apikey)

	// Encode the apikey
	apikey = utils.EncodeAPIKey(apikey)

	result, err := CompanyOverviewUrl(symbol, apikey)
	assert.NoError(t, err)

	assert.Equal(t, result, expectedUrl)

	_, err = CompanyOverviewUrl(symbol, "")
	assert.Error(t, err)
}
