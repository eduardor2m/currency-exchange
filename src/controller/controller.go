package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/eduardor2m/currency-exchange/src/model"
	"github.com/gorilla/mux"
)

func HandleCurrencyConversion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	amount := 1.0
	if amt, ok := params["amount"]; ok {
		amount = parseAmount(amt)
	}

	fromCurrency := strings.ToUpper(params["from"])
	toCurrency := strings.ToUpper(params["to"])

	result, err := model.ConvertCurrency(amount, fromCurrency, toCurrency)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"from":   fromCurrency,
		"to":     toCurrency,
		"amount": amount,
		"result": result,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func parseAmount(amountStr string) float64 {
	var amount float64
	_, err := fmt.Sscanf(amountStr, "%f", &amount)
	if err != nil {
		return 1.0
	}
	return amount
}
