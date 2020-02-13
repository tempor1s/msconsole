package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(linksCommand)
}

var linksCommand = &cobra.Command{
	Use: "links",
	Short: "Get all links for a class.",
	Long: "Given a class code, this command will give you a list of useful links to a class.",
	Run: func(cmd *cobra.Command, args[]string) {
		// TODO: Implement
		fmt.Println("MSConsole's links command. Coming soon! :)")
	},
}
