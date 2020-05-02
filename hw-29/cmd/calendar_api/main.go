package main

import (
	"log"

	"github.com/Temain/otus-golang/hw-29/cmd/calendar_api/api"
)

func main() {
	if err := api.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
