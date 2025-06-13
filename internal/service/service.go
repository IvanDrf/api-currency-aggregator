package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/IvanDrf/currency-aggregator/internal/models"
)

type Parser interface {
	Parse(currency string) (models.Source, error)
}

var (
	parsers = []Parser{&cbr{}, &binance{}}
)

func GetCurrency(currency string) models.Responce {
	sources := workerPool(currency)
	fmt.Println(sources)
	return models.Responce{
		Currency: currency,
		Agerage:  calculateAverage(sources),

		Sources: sources,
	}
}

func workerPool(currency string) []models.Source {
	sources := make(chan models.Source, len(parsers))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	wg := new(sync.WaitGroup)

	for i := 0; i < len(parsers); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			worker(ctx, parsers[i], currency, sources)
		}(i)

	}

	go func() {
		wg.Wait()
		close(sources)
	}()

	result := make([]models.Source, 0, len(parsers))

	for value := range sources {
		fmt.Println(value)
		result = append(result, value)
	}

	return result
}

func worker(ctx context.Context, parser Parser, currency string, sources chan models.Source) {
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

func calculateAverage(sources []models.Source) float64 {
	if len(sources) == 0 {
		return 0
	}

	var summ float64
	for _, value := range sources {
		summ += value.Rate
	}

	return summ / float64(len(sources))
}
