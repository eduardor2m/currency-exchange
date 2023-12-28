package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eduardor2m/currency-exchange/src/controller"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/convert/{from}/{to}/{amount}", controller.HandleCurrencyConversion).Methods("GET")

	port := ":8080"
	fmt.Printf("Listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
