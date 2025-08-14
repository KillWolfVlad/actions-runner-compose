package docker

import (
	"context"

	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/client"
)

func StartContainer(dockerClient *client.Client, containerConfig *container.Config, hostConfig *container.HostConfig) error {
	if err := pullImage(dockerClient, containerConfig.Image); err != nil {
		return err
	}

	containerCreateResponse, err := dockerClient.ContainerCreate(
		context.Background(),
		containerConfig,
		hostConfig,
		nil,
		nil,
		"",
	)

	if err != nil {
		return err
	}

	if err := dockerClient.ContainerStart(context.Background(), containerCreateResponse.ID, container.StartOptions{}); err != nil {
		return err
	}

	return nil
}
