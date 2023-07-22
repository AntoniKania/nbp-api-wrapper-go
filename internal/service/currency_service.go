package service

import "context"

type CurrencyService interface {
	GetAverageCurrencyRate(ctx context.Context, currency string, startDate string, endDate string) float64
}

type currencyService struct {
}

func NewCurrencyService() CurrencyService {
	return &currencyService{}
}
func (s *currencyService) GetAverageCurrencyRate(ctx context.Context, currency string, startDate string, endDate string) float64 {
	//todo: add call to nbp api to collect data and save request to db
	return 5.3
}
