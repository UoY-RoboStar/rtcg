// Package testlang implements the testing language.
package testlang

import (
	"encoding/json"
	"fmt"
	"io"
)

// Suite is a test suite, with each test having a name.
type Suite map[string]*Node

// ReadSuite reads a test suite from JSON in reader r.
func ReadSuite(r io.Reader) (Suite, error) {
	var suite Suite

	j := json.NewDecoder(r)
	if err := j.Decode(&suite); err != nil {
		return nil, fmt.Errorf("json decoding error for test suite: %w", err)
	}

	return suite, nil
}

// Write pretty-prints a test suite, as JSON, into writer w.
func (s *Suite) Write(w io.Writer) error {
	j := json.NewEncoder(w)
	j.SetIndent("", "\t")

	if err := j.Encode(s); err != nil {
		return fmt.Errorf("json encoding error for test suite: %w", err)
	}

	return nil
}

// Node captures a test node of the form (inc -> evt -> X) or (pass -> evt -> X).
type Node struct {
	ID    NodeID   `json:"id,omitempty"`    // ID is an optional identifier assigned to the node.
	Tests []string `json:"tests,omitempty"` // Tests optionally tags a node with the original tests from which it came.

	Status Status `json:"status"`          // Status is the status at this node.
	Event  *Event `json:"event,omitempty"` // Event is the event, if any, contained in this part of the trace.
	Next   []Node `json:"next,omitempty"`  // Next is the list of possible continuations of this test.
}

// Pass constructs a valid passing test node.
func Pass(event Event, next ...Node) Node {
	return NewNode(StatusPass, &event, next...)
}

// Inc constructs a valid inconclusive test node.
func Inc(event Event, next ...Node) Node {
	return NewNode(StatusInc, &event, next...)
}

// Fail constructs a valid failing test node.
func Fail() Node {
	return NewNode(StatusFail, nil)
}

// NewNode constructs a new node with the given status, event, and next nodes.
//
// The node does not have an assigned ID or test list; set these afterwards if desired.
func NewNode(status Status, event *Event, next ...Node) Node {
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
