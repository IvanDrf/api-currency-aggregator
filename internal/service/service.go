package service

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	usd = "USD"
	eur = "EUR"
)

type Parser interface {
	Parse(req *http.Response, currency string) (Source, error)
}

var (
	urls = []string{
		"https://www.cbr-xml-daily.ru/daily_json.js",
		"https://api.binance.com/api/v3/ticker/price?symbol=USDTRUB",
	}

	parsers = []Parser{&cbr{}, &binance{}}
)

func GetCurrency() {
	sources := make(chan Source, len(urls))

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	wg := new(sync.WaitGroup)

	for i := 0; i < len(urls); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			Worker(ctx, parsers[i], urls[i], sources)
		}(i)
	}

	go func() {
		wg.Wait()
		close(sources)
	}()

	for value := range sources {
		fmt.Println(value)
	}
}

func Worker(ctx context.Context, parser Parser, url string, sources chan Source) {
	select {
	case <-ctx.Done():
		return

	default:
		resp, err := http.Get(url)
		if err != nil {
			return
		}

		res, err := parser.Parse(resp, "USD")
		if err != nil {
			return
		}
		defer resp.Body.Close()

		select {
		case sources <- res:
			return
		case <-ctx.Done():
			return
		}
	}

}
