package utils

import (
	"context"
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"strings"
)

func StartInitContainer(compose DockerCompose, blockChannel chan int) error {
	ctx := context.Background()

	initRequests := testcontainers.ParallelContainerRequest{}
	for _, v := range compose.Services {
		if v.Init == true {
			newService := getRequest(v)
			initRequests = append(initRequests, newService)
		}
	}

	_, err := testcontainers.ParallelContainers(ctx, initRequests, testcontainers.ParallelContainersOptions{})

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

	blockChannel <- 666

	return nil
}

func StartDockerByYml(compose DockerCompose) error {
	blockChannel := make(chan int)

	ctx := context.Background()

	requests := testcontainers.ParallelContainerRequest{}

	for _, v := range compose.Services {
		if v.Init == false {
			newService := getRequest(v)
			requests = append(requests, newService)
		}
	}

	go StartInitContainer(compose, blockChannel)

	<-blockChannel

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

func getRequest(item ContainerItem) testcontainers.GenericContainerRequest {
	port := ""
	if len(item.Ports) == 0 {
		port = "80"
	}
	port = strings.Split(item.Ports[0], ":")[0]
	exposedPorts := port + "/tcp"

	newService := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{

			Image: item.Image,
			ExposedPorts: []string{
				port + "/tcp",
			},
			Name:       item.ContainerName,
			WaitingFor: wait.ForListeningPort(nat.Port(exposedPorts)),
		},
		Started: true,
	}

	return newService
}
