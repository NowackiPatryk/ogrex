package cli

import "fmt"

func validateRunCommandArgs(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing config file argument")
	}

	return nil
}