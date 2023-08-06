package testutil

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go/wait"
	"os"

	"github.com/testcontainers/testcontainers-go"
	"github.com/walkerus/go-wiremock"
	"testing"
)

var wiremockClient *wiremock.Client

func SetupWireMock(t *testing.T) (*wiremock.Client, string) {
	ctx, wmContainer := setupWireMockContainer(t)

	ip, _ := wmContainer.Host(ctx)
	port, _ := wmContainer.MappedPort(ctx, "8443")
	wmAdd := fmt.Sprintf("http://%s:%s", ip, port.Port())

	wiremockClient = wiremock.NewClient(wmAdd)
	defer wiremockClient.Reset()

	return wiremockClient, wmAdd
}

func ReadFile(fileName string) string {
	b, _ := os.ReadFile("../../__files/" + fileName)
	return string(b)
}

func setupWireMockContainer(t *testing.T) (context.Context, testcontainers.Container) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "wiremock/wiremock:latest",
		ExposedPorts: []string{"8443/tcp"},
		Cmd:          []string{"--port", "8443"},
		WaitingFor:   wait.ForLog("port:"),
	}

	wmContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Error(err)
	}

	t.Cleanup(func() {
		if err := wmContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})
	return ctx, wmContainer
}
