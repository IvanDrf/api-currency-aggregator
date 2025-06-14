package models

import "math"

type Source struct {
	Currency string  `json:"-"`
	Source   string  `json:"source"`
	Rate     float64 `json:"rate"`
}

func (s *Source) Round() {
	ratio := math.Pow(10, float64(3))
	s.Rate = math.Round(s.Rate*ratio) / ratio
}
