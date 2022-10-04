package utils

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"log"
	"strings"
)

func StartDockerByYml(compose DockerCompose) error {
	ctx := context.Background()

	requests := testcontainers.ParallelContainerRequest{}

	for _, v := range compose.Services {

		if len(v.Ports) == 0 {
			continue
		}
		port := strings.Split(v.Ports[0], ":")[0]

		newService := testcontainers.GenericContainerRequest{
			ContainerRequest: testcontainers.ContainerRequest{

				Image: v.Image,
				ExposedPorts: []string{
					port + "/tcp",
				},
			},
			Started: true,
		}

		requests = append(requests, newService)
	}

	_, err := testcontainers.ParallelContainers(ctx, requests, testcontainers.ParallelContainersOptions{})

	if err != nil {
		e, ok := err.(testcontainers.ParallelContainersError)
		if !ok {
			log.Fatalf("unknown error: %v", err)
		}

		for _, pe := range e.Errors {
			fmt.Println(pe.Request, pe.Error)
		}
		return err
	}

	//for _, c := range res {
	//	defer c.Terminate(ctx)
	//}
	return nil
}
