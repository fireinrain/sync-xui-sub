package task

import (
	"fmt"
	"time"
)

func CheckSubscribeData() {
	// set the target time to 4 AM
	targetTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 4, 0, 0, 0, time.Now().Location())

	// calculate the duration until the target time
	duration := targetTime.Sub(time.Now())

	// if the target time has already passed for today, add 24 hours to the duration
	if duration < 0 {
		duration = duration + 24*time.Hour
	}

	// wait until the target time
	timer := time.NewTimer(duration)
	defer timer.Stop()

	// loop that runs at 24-hour intervals starting at the target time
	for {
		<-timer.C
		// do something here
		fmt.Println("Running task at 4 AM")
		timer.Reset(24 * time.Hour)
	}
}
