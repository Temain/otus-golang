package api

import (
	"fmt"
	"log"

	"github.com/Temain/otus-golang/hw-26/internal/api"

	"github.com/spf13/cobra"
)

var HttpServerCmd = &cobra.Command{
	Use:   "http_server",
	Short: "run http server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("running http server...")

		err := api.StartHttpServer()
		if err != nil {
			log.Fatalf("http server error: %v", err)
		}
	},
}
