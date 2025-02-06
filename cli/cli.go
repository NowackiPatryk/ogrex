package cli

import (
	"fmt"
	"os"
)

func SetupCli() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
