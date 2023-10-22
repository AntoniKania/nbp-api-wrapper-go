package testutil

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/wiremock/go-wiremock"
	"os"

	"github.com/testcontainers/testcontainers-go"
)

var Wm *WireMock

type WireMock struct {
	Client    *wiremock.Client
	Address   string
	ctx       context.Context
	container testcontainers.Container
}

func SetupWireMock() {
	ctx, wmContainer := setupWireMockContainer()

	ip, _ := wmContainer.Host(ctx)
	port, _ := wmContainer.MappedPort(ctx, "8443")
	wmAdd := fmt.Sprintf("http://%s:%s", ip, port.Port())

	wiremockClient := wiremock.NewClient(wmAdd)
	defer wiremockClient.Reset()

	Wm = &WireMock{
		Client:    wiremockClient,
		Address:   wmAdd,
		ctx:       ctx,
		container: wmContainer,
	}
}

func ReadFile(fileName string) string {
	b, _ := os.ReadFile("../../__files/" + fileName)
	return string(b)
}

func setupWireMockContainer() (context.Context, testcontainers.Container) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "wiremock/wiremock:2.35.0-1-alpine",
		ExposedPorts: []string{"8443/tcp"},
		Cmd:          []string{"--port", "8443"},
		WaitingFor:   wait.ForLog("port:"),
	}

	wmContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		panic(err)
	}

	return ctx, wmContainer
}

func (wm *WireMock) CleanUp() {
	if err := wm.container.Terminate(wm.ctx); err != nil {
		panic(err)
	}
}
