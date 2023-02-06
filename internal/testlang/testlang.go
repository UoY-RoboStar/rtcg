// Package testlang implements the testing language.
package testlang

// Node captures a test node of the form (inc -> evt -> X) or (pass -> evt -> X).
type Node struct {
	Status Status `json:"status"`          // IsPassing is true if the test should pass at this point.
	Event  Event  `json:"event,omitempty"` // Event is the event contained in this part of the trace.
	Next   []Node `json:"next,omitempty"`  // Next is the list of possible continuations of this test.
}

// Pass constructs a valid passing test node.
func Pass(event Event, next ...Node) Node {
	return Node{Status: StatPass, Event: event, Next: next}
}

// Inc constructs a valid inconclusive test node.
func Inc(event Event, next ...Node) Node {
	return Node{Status: StatInc, Event: event, Next: next}
}

// Fail constructs a valid failing test node.
func Fail() Node {
	return Node{Status: StatFail}
}

// Event is the type of events in a trace.
//
// An event with an empty channel name is considered to be absent, which is only well-formed for fail nodes.
type Event struct {
	Channel   string `json:"channel,omitempty"`   // Channel is the name of the channel on which the event is occurring.
	Direction InOut  `json:"direction,omitempty"` // InOut is the direction of the event.
	Value     Value  `json:"value,omitempty"`     // Value is the value, if any, carried by this event.
}

// Value is the lowest common denominator for our encoding of RoboChart/CSP values.
type Value any
