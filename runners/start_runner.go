package runners

import (
	"fmt"

	"github.com/KillWolfVlad/actions-runner-compose/configs"
	"github.com/KillWolfVlad/actions-runner-compose/docker"
	"github.com/go-playground/webhooks/v6/github"
	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/api/types/mount"
	"github.com/moby/moby/client"
)

func startRunner(config configs.Config, dockerClient *client.Client, workflowJob github.WorkflowJobPayload) error {
	repositoryConfig, err := config.FindRepositoryConfig(workflowJob)

	if err != nil {
		return err
	}

	err = docker.StartContainer(
		dockerClient,
		&container.Config{
			Image: config.RunnerImage,
			Env: []string{
				fmt.Sprintf("ACCESS_TOKEN=%s", repositoryConfig.AccessToken),
				fmt.Sprintf("REPO_URL=%s", workflowJob.Repository.HTMLURL),
				"EPHEMERAL=true",
				"DISABLE_AUTO_UPDATE=true",
			},
		},
		&container.HostConfig{
			AutoRemove: true,
			Mounts: []mount.Mount{
				{
					Type:   mount.TypeBind,
					Source: "/var/run/docker.sock",
					Target: "/var/run/docker.sock",
				},
			},
		},
	)

	if err != nil {
		return err
	}

	return nil
}
