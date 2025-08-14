package docker

import (
	"log"

	"github.com/moby/moby/client"
)

func InitDockerClient() *client.Client {
	dockerClient, err := client.NewClientWithOpts(client.WithHost("unix:///var/run/docker.sock"), client.WithAPIVersionNegotiation())

	if err != nil {
		log.Fatalln(err)
	}

	return dockerClient
}
