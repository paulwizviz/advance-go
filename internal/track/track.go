// Package csvutil provides utility types and functions to simplify working with CSV data.
//
// This package includes:
//   - A function to read a CSV file in a goroutine and channel data.
//   - A function to help count number of lines in a CSV file
package track

import "fmt"

// Tracker represents an interface to an implementation
// of a tracker.
type Tracker interface {
	Movement(from, to string)
}

// TrackerFunc represents a function type to
// for tracking.
type TrackerFunc func(string, string)

func (t TrackerFunc) Movement(from, to string) {
	t(from, to)
}

// TrackerHandler represents a handler for tracking
// actions of tings
type TrackerHandler struct {
	Action string
}

func (m TrackerHandler) Movement(from, to string) {
	fmt.Printf("%s from %s to %s\n", m.Action, from, to)
}

// MovementOf is a pluggable function to track things.
// It accepts custom tracker to record movement.
func MovementOf(thing, from, to string, tracker Tracker) {
	fmt.Printf("%s ", thing)
	tracker.Movement(from, to)
}
