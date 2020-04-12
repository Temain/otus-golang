package main

import (
	"log"

	"githb.com/Temain/otus-golang/hw-19/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
