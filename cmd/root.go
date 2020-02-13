package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// Global flags
var Verbose bool

// Root command
var rootCmd = &cobra.Command{
	Use: "ms",
	Short: "The hub for all things MakeSchool.",
	Long: `MSConsole is a tool that helps you with your day to day to day tasks at Make School.
Built with love by Ben and Gary in Go <3
The repo can be found at https://github.com/BenAndGarys/msconsole-go`, 	// TODO Update link once we change repo name
	Run: func(cmd *cobra.Command, args[]string) {
		// Print the usage if no args are passed in :)
		if err := cmd.Usage(); err != nil {
			log.Fatal(err)
		}
	},
}

// Execute a command
func Execute() {
	// Global flags
	// TODO: Put these into separate function or something
	// TODO: Implement verbose mode
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output mode")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
