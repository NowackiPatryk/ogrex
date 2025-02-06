package config

type ServiceConfig struct {
	Url      string   `yaml:"url"`
	Services []string `yaml:"services"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type Config struct {
	Server ServerConfig `yaml:"server"`
	Services map[string]ServiceConfig `yaml:",inline"`
}