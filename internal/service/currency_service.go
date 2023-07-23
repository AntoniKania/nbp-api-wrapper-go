package service

import (
	"context"
	"encoding/json"
	"github.com/AntoniKania/nbp-api-wrapper-go/internal/api/model"
	"io"
	"log"
	"net/http"
	url2 "net/url"
	"path"
)

type CurrencyService interface {
	GetAverageCurrencyRate(ctx context.Context, currency string, startDate string, endDate string) float64
}

type currencyService struct {
}

func NewCurrencyService() CurrencyService {
	return &currencyService{}
}

func (s *currencyService) GetAverageCurrencyRate(ctx context.Context, currency string, startDate string, endDate string) float64 {
	url, _ := buildUrl(currency, startDate, endDate)

	rates := getCurrencyRates(url)

	var sum float64
	for _, rate := range rates {
		sum += rate.Mid
	}

	return sum / float64(len(rates))
}

func getCurrencyRates(u *url2.URL) []model.CurrencyRate {
	request, _ := http.NewRequest("GET", u.String(), nil)
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Set("Accept", "*/*")
	request.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11`)

	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		log.Fatalln(err)
	}

	responseBody, _ := io.ReadAll(resp.Body)
	var nbpResponse model.NbpResponse

	json.Unmarshal(responseBody, &nbpResponse)

	defer resp.Body.Close()
	return nbpResponse.Rates
}

func buildUrl(currency string, startDate string, endDate string) (*url2.URL, error) {
	baseUrl := "http://api.nbp.pl/api/exchangerates/rates/a"
	u, err := url2.Parse(baseUrl)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, currency, startDate, endDate)
	return u, nil
}
