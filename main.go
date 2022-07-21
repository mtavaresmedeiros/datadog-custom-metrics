package main

import (
	"log"
	"time"

	"github.com/mtavaresmedeiros/datadog-custom-metrics/cmd/currenttime"
	"github.com/mtavaresmedeiros/datadog-custom-metrics/cmd/datadogclient"
)

func main() {

	datadogClient := datadogclient.Main()

	var currentTime float64

	for {
		currentTime = currenttime.Main()

		datadogClient.Gauge("currenttime.unix", currentTime, []string{"environment:ae-test"}, 20)

		log.Printf("sent currenttime.unix: %f", currentTime)

		time.Sleep(10 * time.Second)
	}
}
