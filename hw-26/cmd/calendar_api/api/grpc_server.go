package api

import (
	"log"

	"github.com/Temain/otus-golang/hw-26/internal/api"

	"github.com/spf13/cobra"
)

var GrpcServerCmd = &cobra.Command{
	Use:   "grpc_server",
	Short: "run grpc server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("running gRPC server...")

		err := api.StartGrpcServer(configPath)
		if err != nil {
			log.Fatalf("gRPC server error: %v", err)
		}
	},
}
