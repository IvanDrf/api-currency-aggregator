package service

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type binance struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func (b *binance) Parse(resp *http.Response, currency string) (Source, error) {
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
