package models

type Source struct {
	Currency string  `json:"-"`
	Source   string  `json:"source"`
	Rate     float64 `json:"rate"`
}
