package sources

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/IvanDrf/currency-aggregator/internal/models"
)

type Binance struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func (b *Binance) Parse(currency string) (models.Source, error) {
	binanceUrl := "https://api.binance.com/api/v3/ticker/price?symbol="

	switch currency {
	case "USD":
		binanceUrl = binanceUrl + "USDTRUB"

	case "EURUSDT":
		binanceUrl = binanceUrl + "EURUSDT"

	default:
		return models.Source{}, errors.New(fmt.Sprintf("binance is not supporting %s", currency))

	}

	resp, err := http.Get(binanceUrl)
	if err != nil {
		return models.Source{}, errors.New("can't send get req to binance")
	}
	defer resp.Body.Close()

	res := Binance{}

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return models.Source{}, errors.New("can't parse binance json")
	}

	rate, err := strconv.ParseFloat(res.Price, 64)

	if err != nil {
		return models.Source{}, errors.New("bad rate in binance")
	}

	return models.Source{Currency: currency, Source: "binance", Rate: rate}, nil
}
