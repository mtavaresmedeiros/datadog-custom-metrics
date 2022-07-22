package currenttime

import (
	"log"
	"os"
	"time"
)

// function responsible for get the current time in unix format
func Main() float64 {

	log.SetOutput(os.Stdout)

	now := time.Now()
	currentTime := float64(now.Unix())

	log.Print("Current time: ", currentTime)

	return currentTime
}
