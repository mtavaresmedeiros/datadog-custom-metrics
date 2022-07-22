package datadogclient

import (
	"log"
	"os"

	"github.com/DataDog/datadog-go/v5/statsd"
)

// function responsible for creating a new DogStatsD client
func Main() *statsd.Client {

	log.SetOutput(os.Stdout)

	statsd, err := statsd.New(os.Getenv("DOGSTATSD_HOST_IP") + ":8125")
	if err != nil {
		log.Fatal("Cannot get a DogStatsD client.")
	}

	return (statsd)
}
