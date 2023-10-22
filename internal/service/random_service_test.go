package service

import (
	. "github.com/AntoniKania/nbp-api-wrapper-go/tests/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/wiremock/go-wiremock"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	SetupWireMock()

	code := m.Run()

	Wm.CleanUp()
	os.Exit(code)
}

func TestGetRandomRequest(t *testing.T) {
	service := NewRandomService(Wm.Address)

	Wm.Client.StubFor(wiremock.Get(wiremock.URLPathEqualTo("/hello")).
		WillReturnResponse(wiremock.NewResponse().
			WithStatus(200).
			WithBody("hello")),
	)

	result := service.GetRandomRequest()

	assert.Equal(t, "hello", result)
}
