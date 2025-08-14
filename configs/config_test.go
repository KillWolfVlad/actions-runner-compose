package configs

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv("PORT", "8080")
	os.Setenv("WEBHOOK_PATH", "/webhooks")
	os.Setenv("WEBHOOK_SECRET", "webhook_secret")
	os.Setenv("MAX_RUNNERS", "-1")
	os.Setenv("RUNNER_IMAGE", "myoung34/github-runner:2.327.1-ubuntu-noble")
	os.Setenv("REPOSITORY_0", "User1/Repo1;User1AccessToken")
	os.Setenv("REPOSITORY_1", "User2/Repo2;User2AccessToken")

	config := LoadConfig()

	if config.Port != 8080 {
		t.Errorf("Port = %d; want %d", config.Port, 8080)
	}

	if config.WebhookPath != "/webhooks" {
		t.Errorf("WebhookPath = %s; want %s", config.WebhookPath, "/webhooks")
	}

	if config.WebhookSecret != "webhook_secret" {
		t.Errorf("WebhookSecret = %s; want %s", config.WebhookSecret, "webhook_secret")
	}

	if config.MaxRunners != -1 {
		t.Errorf("MaxRunners = %d; want %d", config.MaxRunners, -1)
	}

	if config.RunnerImage != "myoung34/github-runner:2.327.1-ubuntu-noble" {
		t.Errorf("RunnerImage = %s; want %s", config.RunnerImage, "myoung34/github-runner:2.327.1-ubuntu-noble")
	}

	if len(config.RepositoryConfigs) != 2 {
		t.Errorf("len(repositoryConfigs) = %d; want 2", len(config.RepositoryConfigs))
	}

	compareRepositoryConfig(t, config.RepositoryConfigs[0], RepositoryConfig{
		FullName:    "User1/Repo1",
		AccessToken: "User1AccessToken",
	})

	compareRepositoryConfig(t, config.RepositoryConfigs[1], RepositoryConfig{
		FullName:    "User2/Repo2",
		AccessToken: "User2AccessToken",
	})
}

func compareRepositoryConfig(t *testing.T, actual RepositoryConfig, expected RepositoryConfig) {
	if actual.FullName != expected.FullName {
		t.Errorf("FullName = %s; want %s", actual.FullName, expected.FullName)
	}

	if actual.AccessToken != expected.AccessToken {
		t.Errorf("AccessToken = %s; want %s", actual.AccessToken, expected.AccessToken)
	}
}
