package rate

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ExchangeRateResponse struct {
	Rates map[string]float64 `json:"rates"`
	Base  string             `json:"base"`
	Date  string             `json:"date"`
}

// 通貨の種類
type Currency string

const (
	JPY Currency = "JPY"
	USD Currency = "USD"
)

// 通貨の種類を指定して、通貨のレートを取得
func CallExchangeRateAPI(callCurrency Currency, specifyCurrency Currency) (float64, error) {
	// 通貨のレートを取得
	url := fmt.Sprintf("https://api.exchangerate-api.com/v4/latest/%s", callCurrency)
	res, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to fetch data: %s", res.Status)
	}

	// レスポンスをデコード
	var exchangeRateResponse ExchangeRateResponse
	err = json.NewDecoder(res.Body).Decode(&exchangeRateResponse)
	if err != nil {
		return 0, err
	}

	// 指定した通貨のレートを取得
	rate, exists := exchangeRateResponse.Rates[string(specifyCurrency)]
	if !exists {
		return 0, fmt.Errorf("%s rate not found", specifyCurrency)
	}

	return rate, nil
}
