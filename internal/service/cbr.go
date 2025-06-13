package service

import (
	"encoding/json"
	"errors"
	"net/http"
)

const cbrUrl = "https://www.cbr-xml-daily.ru/daily_json.js"

type cbr struct {
	Valute map[string]map[string]interface{} `json:"Valute"`
}

func (c *cbr) Parse(currency string) (Source, error) {
	resp, err := http.Get(cbrUrl)
	if err != nil {
		return Source{}, errors.New("can't send get req to cbr")
	}
	defer resp.Body.Close()

	res := cbr{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return Source{}, errors.New("can't parse cbr json")
	}

	rate, ok := res.Valute[currency]["Value"].(float64)
	if !ok {
		return Source{}, errors.New("can't get valutes value")
	}

	return Source{Source: "cbr", Rate: rate}, nil
}
