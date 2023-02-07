// Package testlang implements the testing language.
package testlang

// Suite is a test suite, with each test having a name.
type Suite map[string]*Node

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
