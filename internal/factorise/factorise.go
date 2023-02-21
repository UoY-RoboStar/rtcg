// Package factorise contains the test factorisation algorithm.
package factorise

import (
	"github.com/UoY-RoboStar/rtcg/internal/testlang"
	"github.com/UoY-RoboStar/rtcg/internal/trace"
)

// SpineNodes constructs a list of all 'spine' test nodes reachable in tests from trace after.
//
// # A spine node is an inconclusive node that, initially, contains all failures
//
// If we are at the top of the trace, the spine node has no event, and so represents the test root.
func SpineNodes(tests map[string]trace.Forbidden, after trace.Trace) []testlang.Node {
	last := lastEvent(after)

	failures := Failures(tests, after)

	inSpines := make([]testlang.Node, 0, len(failures))

	outputs := make([]testlang.Node, 0, len(failures))

	for _, f := range failures {
		node := f.Node()

		if f.Event.Channel.IsIn() {
			// Each input gets placed on its own copy of the spine directly.
			inSpines = append(inSpines, testlang.Inc(last, node))
		} else {
			// Outputs get merged into a single spine node.
			outputs = append(outputs, node)
		}
	}

	return append(inSpines, testlang.Inc(last, outputs...))
}

func lastEvent(after trace.Trace) testlang.Event {
	nAfter := len(after)
	if nAfter == 0 {
		return testlang.Event{}
	}

	return after[nAfter-1]
}

// Failures finds, for every forbidden-trace test in tests, the failures reached by following trace after.
func Failures(tests map[string]trace.Forbidden, after trace.Trace) []Failure {
	// Optimistically assume that all tests have failures for the capacity.
	failures := make([]Failure, 0, len(tests))

	for k, t := range tests {
		if t.Prefix.Equals(after) {
			failures = append(failures, Failure{Test: k, Event: t.Forbid})
		}
	}

	return failures
}

// Failure captures an event that causes the failure of a test.
type Failure struct {
	Test  string         // Test is the name of the failed test.
	Event testlang.Event // Event is the event that causes the failure.
}

// Node constructs the 'pass, event, fail' node represented by this Failure.
func (f Failure) Node() testlang.Node {
	fail := testlang.Fail().From(f.Test)
	return testlang.Pass(f.Event, fail).From(f.Test)
}
