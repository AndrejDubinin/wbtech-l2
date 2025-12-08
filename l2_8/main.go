package main

import (
	"fmt"
	"log"

	"github.com/beevik/ntp"
)

const ntpServer = "pool.ntp.org"

func main() {
	ntpTime, err := ntp.Time(ntpServer)
	if err != nil {
		log.Fatal("failed to get NTP time", err)
	}

	fmt.Println("Current time:", ntpTime.Format("2006-01-02 15:04:05.000 MST"))
}
