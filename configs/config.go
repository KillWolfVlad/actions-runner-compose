package configs

import (
	"fmt"
	"log"
	"slices"

	"github.com/go-playground/webhooks/v6/github"
	"github.com/joho/godotenv"
)

type Config struct {
	Port              int
	WebhookPath       string
	WebhookSecret     string
	MaxRunners        int
	RunnerImage       string
	RepositoryConfigs []RepositoryConfig
}

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}

	return Config{
		Port:              getIntEnv("PORT"),
		WebhookPath:       getStringEnv("WEBHOOK_PATH"),
		WebhookSecret:     getStringEnv("WEBHOOK_SECRET"),
		MaxRunners:        getIntEnv("MAX_RUNNERS"),
		RunnerImage:       getStringEnv("RUNNER_IMAGE"),
		RepositoryConfigs: loadRepositoryConfigs(),
	}
}

func (config Config) FindRepositoryConfig(workflowJob github.WorkflowJobPayload) (RepositoryConfig, error) {
	i := slices.IndexFunc(config.RepositoryConfigs, func(r RepositoryConfig) bool {
		return r.FullName == workflowJob.Repository.FullName
	})

	if i == -1 {
		return RepositoryConfig{}, fmt.Errorf("config for repository %s not found", workflowJob.Repository.FullName)
	}

	return config.RepositoryConfigs[i], nil
}
