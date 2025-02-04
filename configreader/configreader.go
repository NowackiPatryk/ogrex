package configreader

import (
	"os"

	"gopkg.in/yaml.v3"
)

func ReadYamlConfigFromPath(path string) (config, error) {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return config{}, &CouldNotReadConfigError{}
	}

	var config config
	yaml.Unmarshal(fileContent, &config)

	return config, nil
}