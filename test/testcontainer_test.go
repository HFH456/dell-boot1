package test

import (
	"context"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"testing"
)

func TestWithRedis(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "redis:latest",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}

	req2 := testcontainers.ContainerRequest{
		Image:        "redis:latest",
		ExposedPorts: []string{"6380/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}

	redisC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	redisC2, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req2,
		Started:          true,
	})
	if err != nil {
		t.Error(err)
	}
	defer redisC.Terminate(ctx)
	defer redisC2.Terminate(ctx)
}
