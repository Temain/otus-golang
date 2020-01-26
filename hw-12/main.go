package main

import (
	"log"
	"os"

	"github.com/Temain/otus-golang/hw-12/envdir"
)

func main() {
	args := os.Args[1:]
	dir, cmd, err := envdir.ParseArgs(args)
	if err != nil {
		log.Fatal(err)
	}

	env, err := envdir.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	code := envdir.RunCmd(cmd, env)
	os.Exit(code)
}
