package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use: "ogrex",
}

var exampleCommand = &cobra.Command{
	Use: "example",
	Short: "Example use short",
	Long: "Example use long",
	Run: handleExampleCommand,
}

func init() {
	rootCommand.AddCommand(exampleCommand)
}

func handleExampleCommand(cmd *cobra.Command, args []string) {
	fmt.Println("Example command received with args: ", args)
}

func SetupCli() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}