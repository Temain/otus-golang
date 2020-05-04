package main

import (
	"log"

	"github.com/spf13/pflag"

	"github.com/Temain/otus-golang/project/internal/api"
)

var configPath string

func init() {
	pflag.StringVarP(&configPath, "config", "c", "configs/config.json", "Config file path")
	pflag.Parse()
}
func main() {
	err := api.StartGrpcServer(configPath)
	if err != nil {
		log.Fatalf("gRPC server error: %v", err)
	}
}
