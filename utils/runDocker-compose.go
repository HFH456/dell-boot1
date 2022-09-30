package utils

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/testcontainers/testcontainers-go"
	"strings"
)

func RunDockerCompose() error {
	composeFilePaths := []string{"./docker-compose.yml"}
	identifier := strings.ToLower(uuid.New().String())

	compose := testcontainers.NewLocalDockerCompose(composeFilePaths, identifier)
	execError := compose.
		WithCommand([]string{"up", "-d"}).
		WithEnv(map[string]string{
			"key1": "value1",
			"key2": "value2",
		}).
		Invoke()
	err := execError.Error
	if err != nil {
		return fmt.Errorf("Could not run compose file: %v - %v", composeFilePaths, err)
	}

	return nil
}

func DownDockerCompose(identifier string) error {
	composeFilePaths := []string{"./docker-compose.yml"}

	compose := testcontainers.NewLocalDockerCompose(composeFilePaths, identifier)
	execError := compose.Down()
	err := execError.Error
	if err != nil {
		return fmt.Errorf("Could not run compose file: %v - %v", composeFilePaths, err)
	}
	return nil
}
