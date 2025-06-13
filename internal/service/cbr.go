package service

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/IvanDrf/currency-aggregator/internal/models"
)

const cbrUrl = "https://www.cbr-xml-daily.ru/daily_json.js"

type cbr struct {
	Valute map[string]map[string]interface{} `json:"Valute"`
}

func (c *cbr) Parse(currency string) (models.Source, error) {
	resp, err := http.Get(cbrUrl)
	if err != nil {
		return models.Source{}, errors.New("can't send get req to cbr")
	}
	defer resp.Body.Close()

	res := cbr{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return models.Source{}, errors.New("can't parse cbr json")
	}

	rate, ok := res.Valute[currency]["Value"].(float64)
	if !ok {
		return models.Source{}, errors.New("can't get valutes value")
	}

	return models.Source{Currency: currency, Source: "cbr", Rate: rate}, nil
}
