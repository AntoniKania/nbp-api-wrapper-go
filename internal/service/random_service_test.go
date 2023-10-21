package service

import (
	"github.com/AntoniKania/nbp-api-wrapper-go/tests/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/wiremock/go-wiremock"
	"testing"
)

func TestGetRandomRequest(t *testing.T) {
	wm, wmAdd := testutil.SetupWireMock(t)

	service := NewRandomService(wmAdd)

	wm.StubFor(wiremock.Get(wiremock.URLPathEqualTo("/hello")).
		WillReturnResponse(wiremock.NewResponse().
			WithStatus(200).
			WithBody("hello")),
	)

	result := service.GetRandomRequest()

	assert.Equal(t, "hello", result)
}
