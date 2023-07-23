package model

type CurrencyRate struct {
	No   string  `json:"no"`
	Date string  `json:"effectiveDate"`
	Mid  float64 `json:"mid"`
}
