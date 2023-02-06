package stocks

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type Stock struct {
	Symbol                     string  `json:",omitempty"`
	AssetType                  string  `json:",omitempty"`
	Name                       string  `json:",omitempty"`
	Description                string  `json:",omitempty"`
	Exchange                   string  `json:",omitempty"`
	Currency                   string  `json:",omitempty"`
	Country                    string  `json:",omitempty"`
	Sector                     string  `json:",omitempty"`
	FiscalYearEnd              string  `json:",omitempty"`
	LatestQuarter              string  `json:",omitempty"`
	MarketCap                  uint64  `json:"MarketCapitalization,omitempty,string"`
	EBITDA                     uint64  `json:",omitempty,string"`
	PERatio                    float32 `json:",omitempty,string"`
	PEGRatio                   float32 `json:",omitempty,string"`
	BookValue                  float32 `json:",omitempty,string"`
	DividendPerShare           float32 `json:",omitempty,string"`
	DividendYield              float32 `json:",omitempty,string"`
	EPS                        float32 `json:",omitempty,string"`
	RevenuePerShareTTM         float32 `json:",omitempty,string"`
	ProfitMargin               float32 `json:",omitempty,string"`
	OperatingMarginTTM         float32 `json:",omitempty,string"`
	ReturnOnAssetsTTM          float32 `json:",omitempty,string"`
	ReturnOnEquityTTM          float32 `json:",omitempty,string"`
	RevenueTTM                 uint64  `json:",omitempty,string"`
	GrossProfitTTM             uint64  `json:",omitempty,string"`
	DilutedEPSTTM              float32 `json:",omitempty,string"`
	QuarterlyEarningsGrowthYOY float32 `json:",omitempty,string"`
	QuarterlyRevenueGrowthYOY  float32 `json:",omitempty,string"`
	AnalystTargetPrice         float32 `json:",omitempty,string"`
	TrailingPE                 float32 `json:",omitempty,string"`
	ForwardPE                  float32 `json:",omitempty,string"`
	PriceToSalesRatioTTM       float32 `json:",omitempty,string"`
	PriceToBookRatio           float32 `json:",omitempty,string"`
	EVToRevenue                float32 `json:",omitempty,string"`
	EVToEBITDA                 float32 `json:",omitempty,string"`
	Beta                       float32 `json:",omitempty,string"`
	WeekHigh52                 float32 `json:"52WeekHigh,omitempty,string"`
	WeekLow52                  float32 `json:"52WeekLow,omitempty,string"`
	DayMovingAverage50         float32 `json:"50DayMovingAverage,omitempty,string"`
	DayMovingAverage200        float32 `json:"200DayMovingAverage,omitempty,string"`
	SharesOutstanding          uint64  `json:",omitempty,string"`
	DividendDate               string  `json:",omitempty"`
	ExDividendDate             string  `json:",omitempty"`
}

func ParseStockOverview(stockJson []byte) (Stock, error) {
	logger := log.WithField("function", "stocks.ParseStockOverview()")
	var stock Stock

	err := json.Unmarshal(stockJson, &stock)
	if err != nil {
		logText := fmt.Sprintf("failed to unmarshall json into stock structure %s", stockJson)
		logger.WithError(err).Error(logText)

		return Stock{}, fmt.Errorf("%s: %w", logText, err)
	}

	return stock, nil
}
