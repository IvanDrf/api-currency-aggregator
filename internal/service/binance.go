package service

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

// USDTRUB
var binanceUrl = "https://api.binance.com/api/v3/ticker/price?symbol="

type binance struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func (b *binance) Parse(currency string) (Source, error) {
	switch currency {
	case "USD":
		binanceUrl = binanceUrl + "USDTRUB"

	case "EURUSDT":
		binanceUrl = binanceUrl + "EURUSDT"

	default:
		return Source{Source: "binance", Rate: 0}, nil

	}

	resp, err := http.Get(binanceUrl)
	if err != nil {
		return Source{}, errors.New("can't send get req to binance")
	}
	defer resp.Body.Close()

	res := binance{}

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return Source{}, errors.New("can't parse binance json")
	}

	rate, err := strconv.ParseFloat(res.Price, 64)

	if err != nil {
		return Source{}, errors.New("bad rate in binance")
	}

	return Source{Source: "binance", Rate: rate}, nil
}
