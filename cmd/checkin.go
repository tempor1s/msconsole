package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tempor1s/msconsole-go/modules"
)

func init() {
	rootCmd.AddCommand(checkinCommand)
}

var checkinCommand = &cobra.Command{
	Use:   "checkin [checkin code]",
	Short: "Check you into your class.",
	Long:  "This command will allow the user to checkin to a class using a code.",
	Run: func(cmd *cobra.Command, args []string) {
		modules.CheckinModule(cmd, args)
	},
}
