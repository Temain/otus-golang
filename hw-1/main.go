package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func main() {
	ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	currentTime := time.Now()
	fmt.Println("Current Time:")
	fmt.Println(currentTime)

	fmt.Println("Exact Time:")
	fmt.Println(ntpTime)
}
