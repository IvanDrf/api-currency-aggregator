package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"strings"
)

const (
	usd = "USD"
	eur = "EUR"

	cbr     = "cbr"
	binance = "binance"
)

var urlsSources = []string{
	"https://www.cbr-xml-daily.ru/daily_json.js",
	"https://api.binance.com/api/v3/ticker/price?symbol=USDTRUB",
}

type Parser interface {
	Parse(req *http.Response, currency string) (Source, error)
}

type cbrResponse struct {
	Valute map[string]map[string]interface{} `json:"Valute"`
}

type binanceResponce struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func GetCurrency() {
	urls := make(chan string, len(urlsSources))
	sources := make(chan Source, len(urlsSources))

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()

	for i := 0; i < len(urlsSources); i++ {
		go currencyWorker(ctx, "USD", urls, sources)
	}

	for i := 0; i < len(urlsSources); i++ {
		urls <- urlsSources[i]
	}
	close(urls)

	for i := 0; i < len(urlsSources); i++ {
		fmt.Println(<-sources)
	}

}

func currencyWorker(ctx context.Context, currency string, urls chan string, sources chan Source) {
	for {
		select {
		case <-ctx.Done():
			return

		default:
			url, ok := <-urls
			if !ok {
				return
			}

			res, err := getCurrency(url, currency)
			if err != nil {
				return
			}

			sources <- res

		}
	}
}

func getCurrency(url, currency string) (Source, error) {
	resp, err := http.Get(url)
	if err != nil {
		return Source{}, err
	}

	defer resp.Body.Close()

	switch {
	case strings.Count(url, cbr) != 0:
		res, err := parseCBR(resp, currency)
		if err != nil {
			return Source{}, err
		}

		return res, nil

	case strings.Count(url, binance) != 0:
		res, err := parseBinance(resp, currency)
		if err != nil {
			return Source{}, nil
		}

		return res, nil

	default:
		return Source{}, errors.New("invalid url")
	}
}

func parseCBR(resp *http.Response, currency string) (Source, error) {
	res := cbrResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return Source{}, errors.New("can't parse cbr json")
	}

	rate, ok := res.Valute[currency]["Value"].(float64)
	if !ok {
		return Source{}, errors.New("can't get valutes value")
	}

	return Source{Source: cbr, Rate: rate}, nil
}

func parseBinance(resp *http.Response, currency string) (Source, error) {
	res := binanceResponce{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return Source{}, errors.New("can't parse binance json")
	}

	rate, err := strconv.ParseFloat(res.Price, 64)
	if err != nil {
		return Source{}, errors.New("bad rate in binance")
	}

	return Source{Source: binance, Rate: rate}, nil
}
