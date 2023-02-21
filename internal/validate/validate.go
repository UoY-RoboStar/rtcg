// Package validate contains validators for rtcg tests.
package validate

import (
	"errors"
	"fmt"
	"github.com/UoY-RoboStar/rtcg/internal/testlang"
)

// Test performs a full validation of the test starting at root.
func Test(root *testlang.Node) error {
	// TODO: check test monotonicity and partitioning

	return testlang.Walk(root, func(node *testlang.Node) error {
		var err error

		if node == root {
			err = Root(node)
		} else {
			err = Node(node)
		}

		if err != nil {
			return fmt.Errorf("invalid node %q: %w", node, err)
		}

		return nil
	})
}

// Root checks whether node is a valid root node.
func Root(node *testlang.Node) error {
	if node.Event != nil {
		return ErrRootHasEvent
	}

	if node.Outcome != testlang.OutcomeUnset {
		return ErrRootHasOutcome
	}

	if len(node.Tests) == 0 {
		return ErrNoTests
	}

	return nil
}

// Node checks whether node is a valid non-root node.
func Node(node *testlang.Node) error {
	if len(node.Tests) == 0 {
		return ErrNoTests
	}

	switch node.Outcome {
	case testlang.OutcomeUnset:
		return ErrNoOutcome
	case testlang.OutcomeFail:
		return failNode(node)
	case testlang.OutcomePass:
		if err := passNode(node); err != nil {
			return err
		}
		fallthrough
	case testlang.OutcomeInc:
		return nonFailNode(node)
	default:
		return fmt.Errorf("%w: %s", ErrBadOutcome, node.Outcome)
	}
}

// failNode checks whether node is a valid failing node.
func failNode(node *testlang.Node) error {
	if node.Event != nil {
		return ErrFailHasEvent
	}

	if len(node.Next) != 0 {
		return ErrFailHasNextNodes
	}

	return nil
}

// passNode checks whether node is a valid passing node (on top of the checks in nonFailNode).
func passNode(node *testlang.Node) error {
	if _, err := OneNext(node); err != nil {
		return fmt.Errorf("%w (passing nodes must have exactly one next node)", err)
	}

	return nil
}

// nonFailNode checks whether node is a valid non-failing node.
func nonFailNode(node *testlang.Node) error {
	if node.Event == nil {
		return ErrNoEvent
	}

	if len(node.Next) == 0 {
		return ErrNoNextNodes
	}

	return nil
}

// OneNext tries to assert that there is exactly one next node in node, and gets it if successful.
//
// Passing nodes need to have exactly one next node (the failing node), but we export this
// function as it is useful for collapsing trace tests.
func OneNext(node *testlang.Node) (*testlang.Node, error) {
	if nnodes := len(node.Next); nnodes != 1 {
		return nil, fmt.Errorf("%w: node count was %d", ErrNeedOneNode, nnodes)
	}

	return &node.Next[0], nil
}

var (
	// ErrBadOutcome occurs when a non-root node has an out-of-range outcome.
	ErrBadOutcome = errors.New("non-root node has an invalid outcome set")

	// ErrFailHasEvent occurs when a failing node has an event.
	ErrFailHasEvent = errors.New("failing node should not have an event")

	// ErrFailHasNextNodes occurs when a failing node has next nodes.
	ErrFailHasNextNodes = errors.New("failing node should not have next tests")

	// ErrNeedOneNode occurs when a node doesn't have one next node, but should.
	ErrNeedOneNode = errors.New("expected this node to have exactly one next node")

	// ErrNoEvent occurs when a non-failing non-root node has no event.
	ErrNoEvent = errors.New("non-root, non-failing node should have an event set")

	// ErrNoNextNodes occurs when a non-failing node has no onwards nodes.
	ErrNoNextNodes = errors.New("non-failing node should have at least one next node")

	// ErrNoOutcome occurs when a non-root node has no outcome.
	ErrNoOutcome = errors.New("non-root node should have an outcome set")

	// ErrNoTests occurs when a node has no tests set.
	ErrNoTests = errors.New("node should belong to at least one test")

	// ErrRootHasEvent occurs when a root node has an event set.
	ErrRootHasEvent = errors.New("root should not have an event set")

	// ErrRootHasOutcome occurs when a root node has an outcome set.
	ErrRootHasOutcome = errors.New("root should not have an outcome set")
)
