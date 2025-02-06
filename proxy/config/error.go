package config

type couldNotReadConfigError struct {}

func (err *couldNotReadConfigError) Error() string {
	return "Could not read .yaml config file"
}

var CouldNotReadConfigError = couldNotReadConfigError{}