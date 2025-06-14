package sources

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/IvanDrf/currency-aggregator/internal/models"
)

const coinUrl = "https://api.coingecko.com/api/v3/exchange_rates"

type Coingecko struct {
	Valute map[string]map[string]interface{} `json:"rates"`
}

func (c *Coingecko) Parse(currency string) (models.Source, error) {
	resp, err := http.Get(coinUrl)
	if err != nil {
		return models.Source{}, errors.New("can't send get to coingecko")
	}
	defer resp.Body.Close()

	res := Coingecko{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return models.Source{}, errors.New("can't parse coingecko json")
	}

	rate, ok := res.Valute[strings.ToLower(currency)]["value"].(float64)
	if !ok {
		return models.Source{}, errors.New("can't get valutes value")
	}

	rubs, ok := res.Valute["rub"]["value"].(float64)
	if !ok {
		return models.Source{}, errors.New("cant get valutes value")
	}

	return models.Source{Currency: currency, Source: "coingecko", Rate: rubs / rate}, nil
}
