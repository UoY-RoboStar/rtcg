package validate

import (
	"fmt"
	"github.com/UoY-RoboStar/rtcg/internal/testlang"
)

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
