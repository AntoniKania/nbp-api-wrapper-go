package model

type NbpResponse struct {
	Table    string         `json:"table"`
	Currency string         `json:"currency"`
	Code     string         `json:"code"`
	Rates    []CurrencyRate `json:"rates"`
}
