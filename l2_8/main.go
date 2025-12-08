package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

const ntpServer = "pool.ntp.org"

func main() {
	ntpTime, err := ntp.Time(ntpServer)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println("Current time:", ntpTime.Format("2006-01-02 15:04:05.000 MST"))
}
