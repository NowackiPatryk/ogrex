package configreader

type serviceConfig struct {
	Url      string   `yaml:"url"`
	Services []string `yaml:"services"`
}

type config struct {
	Server struct {
		port int
	} `yaml:"server"`

	Services map[string]serviceConfig `yaml:",inline"`
}