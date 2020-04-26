package api

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "calendar",
	Short: "Simple calendar",
}

func init() {
	RootCmd.AddCommand(HttpServerCmd)
	RootCmd.AddCommand(GrpcServerCmd)
	RootCmd.AddCommand(GrpcClientCmd)
}
