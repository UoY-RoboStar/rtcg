// Package testlang implements the testing language.
package testlang

// TestNode captures a test node of the form (inc -> evt -> X) or (pass -> evt -> X).
type TestNode struct {
	Status Status     // IsPassing is true if the test should pass at this point.
	Event  Event      // Event is the event contained in this part of the trace.
	Next   []TestNode // Next is the list of possible continuations of this test.
}

// Pass constructs a valid passing test node.
func Pass(event Event, next ...TestNode) TestNode {
	return TestNode{Status: StatPass, Event: event, Next: next}
}

// Inc constructs a valid inconclusive test node.
func Inc(event Event, next ...TestNode) TestNode {
	return TestNode{Status: StatInc, Event: event, Next: next}
}

// Fail constructs a valid failing test node.
func Fail() TestNode {
	return TestNode{Status: StatFail}
}

// Status is an enumeration of statuses during a test.
type Status uint8

const (
	StatInc  Status = iota // StatInc is the inconclusive status.
	StatFail               // StatFail is the failing status.
	StatPass               // StatPass is the passing status.
)

// Event is the type of events in a trace.
//
// An event with an empty channel name is considered to be absent, which is only well-formed for fail nodes.
type Event struct {
	Channel   string // Channel is the name of the channel on which the event is occurring.
	Direction InOut  // InOut is the direction of the event.
	Data      Value
}

// Value is the lowest common denominator for our encoding of RoboChart/CSP values.
type Value any

// InOut is an enumeration of communication directions.
type InOut bool

const (
	In  InOut = false // In represents an input.
	Out InOut = true  // Out represents output.
)
