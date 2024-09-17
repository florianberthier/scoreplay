package service

import (
	"context"
	"scoreplay/query"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func (s *Service) CleanupMock() {
	s.DB.DB.Exec("TRUNCATE TABLE tags RESTART IDENTITY CASCADE")
	s.DB.DB.Exec("TRUNCATE TABLE media RESTART IDENTITY CASCADE")
}

func SetupMock(t *testing.T) *Service {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "postgres:14-alpine",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_PASSWORD": "password",
			"POSTGRES_USER":     "user",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}

	postgresContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatalf("Could not start container: %s", err)
	}

	host, err := postgresContainer.Host(ctx)
	if err != nil {
		t.Fatalf("Could not get host: %s", err)
	}

	port, err := postgresContainer.MappedPort(ctx, "5432")
	if err != nil {
		t.Fatalf("Could not get port: %s", err)
	}

	return &Service{
		DB:        query.Setup(host, port.Port(), "user", "password", "testdb"),
		Validator: validator.New(),
	}
}
