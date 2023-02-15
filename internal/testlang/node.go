package testlang

// Node captures a test node of the form (inc -> evt -> X), (pass -> evt -> X), or (fail).
type Node struct {
	ID    NodeID   `json:"id,omitempty"`    // ID is an optional identifier assigned to the node.
	Tests []string `json:"tests,omitempty"` // Tests optionally tags a node with the original tests from which it came.

	Status Outcome `json:"status"`          // Status is the status at this node.
	Event  *Event  `json:"event,omitempty"` // Event is the event, if any, contained in this part of the trace.
	Next   []Node  `json:"next,omitempty"`  // Next is the list of possible continuations of this test.
}

// Pass constructs a valid passing test node.
func Pass(event Event, next ...Node) Node {
	return NewNode(OutcomePass, &event, next...)
}

// Inc constructs a valid inconclusive test node.
func Inc(event Event, next ...Node) Node {
	return NewNode(OutcomeInc, &event, next...)
}

// Fail constructs a valid failing test node.
func Fail() Node {
	return NewNode(OutcomeFail, nil)
}

// NewNode constructs a new node with the given status, event, and next nodes.
//
// The node does not have an assigned ID or test list; set these afterwards if desired.
func NewNode(status Outcome, event *Event, next ...Node) Node {
	return Node{ID: "", Tests: nil, Status: status, Event: event, Next: next}
}

// From replaces the Tests field of this Node inline with the contents of tests.
func (n Node) From(tests ...string) Node {
	n.Tests = tests

	return n
}

// NodeID is the type of node identifiers that can be attached to a node.
//
// Generally, these will be added by the generator.
type NodeID string
