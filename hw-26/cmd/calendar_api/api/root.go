package api

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "calendar",
	Short: "Simple calendar",
}

var configPath string

func init() {
	RootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "configs/config.json", "Config file path")

	RootCmd.AddCommand(HttpServerCmd)
	RootCmd.AddCommand(GrpcServerCmd)
	RootCmd.AddCommand(GrpcClientCmd)
}
