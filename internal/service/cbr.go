package service

import (
	"encoding/json"
	"errors"
	"net/http"
)

type cbr struct {
	Valute map[string]map[string]interface{} `json:"Valute"`
}

func (c *cbr) Parse(resp *http.Response, currency string) (Source, error) {
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
