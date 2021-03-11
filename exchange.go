package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Money float64
type Exchange struct {
	Rates struct {
		BTC float64 `json:"BTC"`
		EUR float64 `json:"EUR"`
		KZT float64 `json:"KZT"`
		RUB float64 `json:"RUB"`
		USD int     `json:"USD"`
	}
}

var (
	bitcoin string
	Euro    string
	Tenge   string
	Dollar  string
)

func exchange() {
	rates := Exchange{}
	urll := fmt.Sprintf("https://openexchangerates.org/api/latest.json?app_id=%s", cfg.ExchangeID)
	resps, err := http.Get(urll)
	if err != nil {
		log.Fatal(err)

	}
	defer resps.Body.Close()
	json.NewDecoder(resps.Body).Decode(&rates)

	Euro = fmt.Sprintf("%.2f руб.", 1/rates.Rates.EUR*rates.Rates.RUB)
	Dollar = fmt.Sprintf("%.2f руб.", rates.Rates.RUB)
	Tenge = fmt.Sprintf("%.2f руб.", rates.Rates.KZT/rates.Rates.RUB)
	bitcoin = fmt.Sprintf("%.2f руб.", 1/rates.Rates.BTC*rates.Rates.RUB)

	fmt.Println(Euro, Dollar, Tenge, bitcoin)
}
