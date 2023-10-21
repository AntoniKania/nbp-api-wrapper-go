package service

import (
	"io"
	"net/http"
)

type RandomService interface {
	GetRandomRequest() string
}

type randomService struct {
	targetAddress string
}

func NewRandomService(targetAddress string) RandomService {
	return &randomService{
		targetAddress: targetAddress,
	}
}

func (r *randomService) GetRandomRequest() string {
	baseUrl := r.targetAddress + "/hello"

	request, _ := http.NewRequest("GET", baseUrl, nil)

	client := &http.Client{}
	resp, _ := client.Do(request)

	respBody, _ := io.ReadAll(resp.Body)

	return string(respBody)
}
