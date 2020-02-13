package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(checkinCommand)
}

var checkinCommand = &cobra.Command{
	Use: "checkin [checkin code]",
	Short: "Check you into your class.",
	Long: "This command will allow the user to checkin to a class using a code.",
	Run: func(cmd *cobra.Command, args[]string) {
		if len(args) == 0 {
			fmt.Println("Please enter a a checkin code to this command. Example: `ms checkin dog`")
			return
		}

		fmt.Printf("Your checkin code is %s\n", args[0])
	},
}
