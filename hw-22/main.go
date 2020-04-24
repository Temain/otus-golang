package main

import (
	"log"

	"github.com/Temain/otus-golang/hw-22/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
