package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "calendar",
	Short: "Simple calendar",
}

func init() {
	RootCmd.AddCommand(SchedulerCmd)
	RootCmd.AddCommand(SenderCmd)
}
