package model

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const exchangeRatesAPIURL = "https://open.er-api.com/v6/latest"

type ExchangeRatesResponse struct {
	Rates map[string]float64 `json:"rates"`
}

func GetExchangeRates(baseCurrency string) (ExchangeRatesResponse, error) {
	url := fmt.Sprintf("%s?base=%s", exchangeRatesAPIURL, baseCurrency)
	resp, err := http.Get(url)
	if err != nil {
		return ExchangeRatesResponse{}, err
	}
	defer resp.Body.Close()

	var ratesResponse ExchangeRatesResponse
	err = json.NewDecoder(resp.Body).Decode(&ratesResponse)
	if err != nil {
		return ExchangeRatesResponse{}, err
	}

	return ratesResponse, nil
}

func ConvertCurrency(amount float64, fromCurrency string, toCurrency string) (float64, error) {
	ratesResponse, err := GetExchangeRates(fromCurrency)
	if err != nil {
		return 0, err
	}

	rate, exists := ratesResponse.Rates[toCurrency]
	if !exists {
		return 0, fmt.Errorf("no exchange rate found for %s to %s", fromCurrency, toCurrency)
	}

	result := amount * rate
	return result, nil
}
