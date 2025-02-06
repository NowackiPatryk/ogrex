package cli

import (
	"errors"
	"fmt"

	"example.com/ogrex/proxy"
	"example.com/ogrex/proxy/config"
	"github.com/spf13/cobra"
)

func handleRunCommand(cmd *cobra.Command, args []string) {
	if err := validateRunCommandArgs(args); err != nil {
		fmt.Println("Wrong args: ", err.Error())
		return
	}

	path := args[0]

	proxyConfig, err := config.ReadYamlConfigFromPath(path)

	if errors.Is(err, &config.CouldNotReadConfigError) {
		fmt.Println("Config can't be read")
		return
	}

	if err != nil {
		panic("Unknown error occured")
	}

	proxy := proxy.NewProxy(proxyConfig)
	proxy.Run()
}