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
	var collapsed Forbidden

	// We validate the test in-place, as collapsing is effectively a stricter form of validation.
	if err := validate.Root(test); err != nil {
		return collapsed, fmt.Errorf("test does not start with a valid root: %w", err)
	}

	for i := 1; ; i++ {
		var err error
		if test, err = validate.OneNext(test); err != nil {
			return collapsed, fmt.Errorf("couldn't get node %d: %w", i, err)
		}

		if err := validate.Node(test); err != nil {
			return collapsed, fmt.Errorf("node %d failed verification: %w", i, err)
		}

		switch test.Outcome {
		case testlang.OutcomeFail:
			return collapsed, nil
		case testlang.OutcomePass:
			// The next node is the fail node, which we need to inspect, but this one has the
			// forbidden event.
			collapsed.Forbid = *test.Event
		default:
			collapsed.Prefix = append(collapsed.Prefix, *test.Event)
		}
	}
}

// CollapseAll collapses a suite of linear tests into Forbidden traces.
func CollapseAll(suite testlang.Suite) (map[string]Forbidden, error) {
	return structure.TryOverMapValues(suite, Collapse)
}
