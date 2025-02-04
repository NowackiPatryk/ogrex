package configreader

type CouldNotReadConfigError struct {}

func (err *CouldNotReadConfigError) Error() string {
	return "Could not read .yaml config file"
}