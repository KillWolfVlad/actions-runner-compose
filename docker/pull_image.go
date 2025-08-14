package docker

import (
	"context"
	"io"

	"github.com/moby/moby/api/types/image"
	"github.com/moby/moby/client"
)

func pullImage(dockerClient *client.Client, imageName string) error {
	reader, err := dockerClient.ImagePull(context.Background(), imageName, image.PullOptions{})

	if err != nil {
		return err
	}

	defer reader.Close()

	_, err = io.Copy(io.Discard, reader)

	if err != nil {
		return err
	}

	return nil
}
