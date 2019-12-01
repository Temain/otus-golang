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
		fmt.Fprintln(os.Stderr, "error: ", err)
		os.Exit(1)
	}

	currentTime := time.Now()
	fmt.Println("Current Time: ", currentTime)
	fmt.Println("Exact Time: ", ntpTime)
}
