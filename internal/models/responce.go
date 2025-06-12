package models

import "github.com/IvanDrf/currency-aggregator/internal/service"

type Responce struct {
	Currency string           `json:"currency"`
	Agerage  float64          `json:"average_rate"`
	Sources  []service.Source `json:"sources"`

	Data string `json:"data"`
}
