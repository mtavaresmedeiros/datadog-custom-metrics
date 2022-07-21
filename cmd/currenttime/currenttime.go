package currenttime

import (
	"log"
	"os"
	"time"
)

func Main() float64 {

	log.SetOutput(os.Stdout)

	now := time.Now()
	currentTime := float64(now.Unix())

	log.Print("Current time: ", currentTime)

	return currentTime
}
