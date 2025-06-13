package service

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Parser interface {
	Parse(currency string) (Source, error)
}

var (
	parsers    = []Parser{&cbr{}, &binance{}}
	currencies = []string{"USD", "EUR"}
)

func GetCurrency() {
	sources := make(chan Source, len(parsers)*len(currencies))

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	wg := new(sync.WaitGroup)

	for i := 0; i < len(currencies); i++ {
		for j := 0; j < len(currencies); j++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				Worker(ctx, parsers[i], currencies[j], sources)
			}(i)
		}
	}

	go func() {
		wg.Wait()
		close(sources)
	}()

	for value := range sources {
		fmt.Println(value)
	}
}

func Worker(ctx context.Context, parser Parser, currency string, sources chan Source) {
	select {
	case <-ctx.Done():
		return

	default:
		res, err := parser.Parse(currency)
		if err != nil {
			return
		}

		select {
		case sources <- res:
			return
		case <-ctx.Done():
			return
		}
	}

}
