package models

type Responce struct {
	Currency string   `json:"currency"`
	Agerage  float64  `json:"average_rate"`
	Sources  []Source `json:"sources"`
}
