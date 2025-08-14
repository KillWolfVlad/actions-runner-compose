package configs

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type RepositoryConfig struct {
	FullName    string
	AccessToken string
}

func loadRepositoryConfigs() []RepositoryConfig {
	var repositoryConfigs []RepositoryConfig

	for i := 0; ; i++ {
		envKey := fmt.Sprintf("REPOSITORY_%d", i)

		rawConfig, isConfigExists := os.LookupEnv(envKey)

		if !isConfigExists {
			break
		}

		config, err := parseRepositoryConfig(envKey, rawConfig)

		if err != nil {
			log.Fatalln(err)
		}

		repositoryConfigs = append(repositoryConfigs, config)
	}

	return repositoryConfigs
}

func parseRepositoryConfig(envKey string, rawConfig string) (RepositoryConfig, error) {
	parts := strings.Split(rawConfig, ";")

	if len(parts) != 2 {
		return RepositoryConfig{}, fmt.Errorf("%s has invalid format", envKey)
	}

	fullName := strings.TrimSpace(parts[0])
	accessToken := strings.TrimSpace(parts[1])

	if fullName == "" || accessToken == "" {
		return RepositoryConfig{}, fmt.Errorf("%s has invalid format", envKey)
	}

	return RepositoryConfig{
		FullName:    fullName,
		AccessToken: accessToken,
	}, nil
}
