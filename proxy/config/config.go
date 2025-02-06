package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func ReadYamlConfigFromPath(path string) (Config, error) {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return Config{}, &CouldNotReadConfigError
	}

	var config Config
	yaml.Unmarshal(fileContent, &config)

	return config, nil
}
