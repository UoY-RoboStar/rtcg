package trace

import (
	"fmt"

	"github.com/UoY-RoboStar/rtcg/internal/structure"
	"github.com/UoY-RoboStar/rtcg/internal/testlang"
	"github.com/UoY-RoboStar/rtcg/internal/validate"
)

// Collapse tries to perform the inverse of Expand: convert a test tree into a forbidden trace.
// It fails if the tree is ill-formed or has branches.
func Collapse(test *testlang.Node) (Forbidden, error) {
	var (
		collapsed Forbidden
		err       error
	)

	// We validate the test in-place, as collapsing is effectively a stricter form of validation.
	if err = validate.Root(test); err != nil {
		return collapsed, fmt.Errorf("test does not start with a valid root: %w", err)
	}

	for done, pos := false, 1; !done; pos++ {
		if test, err = validate.OneNext(test); err != nil {
			return collapsed, fmt.Errorf("couldn't get node %d: %w", pos, err)
		}

		if err := validate.Node(test); err != nil {
			return collapsed, fmt.Errorf("node %d failed verification: %w", pos, err)
		}

		if done, err = collapseNode(&collapsed, test, pos); err != nil {
			return collapsed, err
		}
	}

	return collapsed, nil
}

func collapseNode(collapsed *Forbidden, test *testlang.Node, pos int) (bool, error) {
	switch test.Outcome {
	case testlang.OutcomeFail:
		return true, nil
	case testlang.OutcomePass:
		// The next node is the fail node, which we need to inspect, but this one has the
		// forbidden event.
		collapsed.Forbid = *test.Event
	case testlang.OutcomeInc:
		collapsed.Prefix = append(collapsed.Prefix, *test.Event)
	case testlang.OutcomeUnset:
		fallthrough
	default:
		return false, BadOutcomeError{Position: pos, Outcome: test.Outcome}
	}

	return false, nil
}

// BadOutcomeError is an error that arises if we see an unwanted outcome while collapsing a test.
type BadOutcomeError struct {
	Position int              // Position is the 1-indexed position where we got the bad outcome.
	Outcome  testlang.Outcome // Outcome is the outcome we saw.
}

func (b BadOutcomeError) Error() string {
	return fmt.Sprintf("unexpected outcome at node %d: %q", b.Position, b.Outcome)
}

// CollapseAll collapses a suite of linear tests into Forbidden traces.
func CollapseAll(suite testlang.Suite) (map[string]Forbidden, error) {
	collapsed, err := structure.TryOverMapValues(suite, Collapse)
	if err != nil {
		return nil, fmt.Errorf("couldn't collapse suite: %w", err)
	}

	return collapsed, nil
}
