package main

import (
	"fmt"
	"go-pattern/internal/track"
)

func main() {
	var walkingTracker track.TrackerFunc = func(from, to string) {
		fmt.Printf("walked from %s to %s\n", from, to)
	}
	track.MovementOf("Alice", "location 1", "location 2", walkingTracker)

	var runningTracker track.TrackerFunc = func(from, to string) {
		fmt.Printf("ran from %s to %s\n", from, to)
	}
	track.MovementOf("Bob", "location 1", "location 2", runningTracker)

	drivingTracker := track.TrackerHandler{
		Action: "drove",
	}
	track.MovementOf("Charlie", "location 1", "location 2", drivingTracker)

	flyingTracker := track.TrackerHandler{
		Action: "flew",
	}
	track.MovementOf("Delta", "location 1", "location 2", flyingTracker)

}
