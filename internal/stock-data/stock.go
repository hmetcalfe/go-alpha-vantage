package stocks

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type Stock struct {
	Symbol                     string  `json:"Symbol"`
	AssetType                  string  `json:"AssetType"`
	Name                       string  `json:"Name"`
	Description                string  `json:"Description"`
	Exchange                   string  `json:"Exchange"`
	Currency                   string  `json:"Currency"`
	Country                    string  `json:"Country"`
	Sector                     string  `json:"Sector"`
	FiscalYearEnd              string  `json:"FiscalYearEnd"`
	LatestQuarter              string  `json:"LatestQuarter"`
	MarketCap                  uint64  `json:"MarketCapitalization"`
	EBITDA                     uint64  `json:"EBITDA"`
	PERatio                    float32 `json:"PERatio"`
	PEGRatio                   float32 `json:"PEGRatio"`
	BookValue                  float32 `json:"BookValue"`
	DividendPerShare           float32 `json:"DividendPerShare"`
	DividendYield              float32 `json:"DividendYield"`
	EPS                        float32 `json:"EPS"`
	RevenuePerShareTTM         float32 `json:"RevenuePerShareTTM"`
	ProfitMargin               float32 `json:"ProfitMargin"`
	OperatingMarginTTM         float32 `json:"OperatingMarginTTM"`
	ReturnOnAssetsTTM          float32 `json:"ReturnOnAssetsTTM"`
	ReturnOnEquityTTM          float32 `json:"ReturnOnEquityTTM"`
	RevenueTTM                 uint64  `json:"RevenueTTM"`
	GrossProfitTTM             uint64  `json:"GrossProfitTTM"`
	DilutedEPSTTM              float32 `json:"DilutedEPSTTM"`
	QuarterlyEarningsGrowthYOY float32 `json:"QuarterlyEarningsGrowthYOY"`
	QuarterlyRevenueGrowthYOY  float32 `json:"QuarterlyRevenueGrowthYOY"`
	AnalystTargetPrice         float32 `json:"AnalystTargetPrice"`
	TrailingPE                 float32 `json:"TrailingPE"`
	ForwardPE                  float32 `json:"ForwardPE"`
	PriceToSalesRatioTTM       float32 `json:"PriceToSalesRatioTTM"`
	PriceToBookRatio           float32 `json:"PriceToBookRatio"`
	EVToRevenue                float32 `json:"EVToRevenue"`
	EVToEBITDA                 float32 `json:"EVToEBITDA"`
	Beta                       float32 `json:"Beta"`
	WeekHigh52                 float32 `json:"52WeekHigh"`
	WeekLow52                  float32 `json:"52WeekLow"`
	DayMovingAverage50         float32 `json:"50DayMovingAverage"`
	DayMovingAverage200        float32 `json:"200DayMovingAverage"`
	SharesOutstanding          uint64  `json:"SharesOutstanding"`
	DividendDate               string  `json:"DividendDate"`
	ExDividendDate             string  `json:"ExDividendDate"`
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
