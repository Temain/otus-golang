package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "telnet",
	Short: "Simple telnet usage example",
}

func init() {
	RootCmd.AddCommand(TelnetServerCmd)
	RootCmd.AddCommand(TelnetClientCmd)
}
