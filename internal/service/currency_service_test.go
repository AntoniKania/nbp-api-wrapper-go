package service

import (
	"context"
	"github.com/AntoniKania/nbp-api-wrapper-go/internal/database/mock_repository"
	"github.com/AntoniKania/nbp-api-wrapper-go/tests/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/wiremock/go-wiremock"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestGetAverageCurrencyRate(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := mock_repository.NewMockRequestRepository(ctrl)
	wm, wmAdd := testutil.SetupWireMock(t)
	service := NewCurrencyService(wmAdd, m)

	wm.StubFor(wiremock.Get(wiremock.URLPathEqualTo("/api/exchangerates/rates/a/gbp/2023-06-15/2023-06-27")).
		WillReturnResponse(wiremock.NewResponse().
			WithStatus(200).
			WithBody(testutil.ReadFile("response.json")).
			WithHeader("Content-Type", "application/json"),
		))

	m.
		EXPECT().
		SaveRequest(gomock.Any(), gomock.Any()).
		Return(int64(1), nil)

	result := service.GetAverageCurrencyRate(context.Background(), "gbp", "2023-06-15", "2023-06-27")

	assert.Equal(t, 5.186711111111112, result)
}
