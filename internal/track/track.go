package track

import "fmt"

// Tracker represents an interface to an implementation
// of a system to track movement of things.
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

// MovementOf is a function to track things moving
func MovementOf(thing, from, to string, tracker Tracker) {
	fmt.Printf("%s ", thing)
	tracker.Movement(from, to)
}
