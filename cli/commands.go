package cli

import "github.com/spf13/cobra"

var rootCommand = &cobra.Command{
	Use: "ogrex",
}

var runCommand = &cobra.Command{
	Use:   "run",
	Short: "Runs with given .yaml config file path",
	Run:   handleRunCommand,
}

func init() {
	rootCommand.AddCommand(runCommand)
}
