package cli

import (
	"fmt"

	"example.com/ogrex/configreader"
	"github.com/spf13/cobra"
)

func handleRunCommand(cmd *cobra.Command, args []string) {
	if err := validateRunCommandArgs(args); err != nil {
		fmt.Println("Wrong args: ", err.Error())
		return
	}

	path := args[0]

	config, err := configreader.ReadYamlConfigFromPath(path)
	if err, ok := err.(*configreader.CouldNotReadConfigError); ok {
		fmt.Println(err.Error())
		return
	}

	if err != nil {
		panic("Unknown error occured")
	}

	fmt.Println("Config successfully loaded: ", config)
}