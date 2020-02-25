package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCommand)
}

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Print the version for MSConsole",
	Long:  "All software has versions. This is MSConsole's",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Pull from somewhere
		fmt.Println("MSConsole's version is currently alpha")
	},
}
