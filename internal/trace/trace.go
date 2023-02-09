// Package trace concerns the ingestion of tests as forbidden trace.
//
// The format of a forbidden trace input file is a CSV file where each line is a separate trace (test) and each
// cell is an event; the last event in
package trace

import (
	"fmt"
	"strings"

	"github.com/UoY-RoboStar/rtcg/internal/testlang"
)

// Trace is the type of 'flat' forbidden-trace tests.
type Trace struct {
	Prefix []testlang.Event // Prefix is the sequence of events that must occur for the test to pass.
	Forbid testlang.Event   // Forbid is the event that must not occur after Prefix.
}

// New constructs a trace with forbidden event forbid and prefix after.
func New(forbid testlang.Event, after ...testlang.Event) Trace {
	return Trace{Prefix: after, Forbid: forbid}
}

func (t Trace) String() string {
	prefixStrs := make([]string, len(t.Prefix))

	for i, p := range t.Prefix {
		prefixStrs[i] = p.String()
	}

	prefixStr := strings.Join(prefixStrs, ", ")

	return fmt.Sprintf("<%s>!%s", prefixStr, &t.Forbid)
}

// Expand expands a single Trace into a test, tagged throughout with name.
func (t Trace) Expand(name string) *testlang.Node {
	// Work backwards through the trace, building the tree from the failure.
	n := testlang.Pass(t.Forbid, testlang.Fail().From(name)).From(name)
	for i := len(t.Prefix) - 1; 0 <= i; i-- {
		n = testlang.Inc(t.Prefix[i], n).From(name)
	}

	return &n
}

// ExpandAll expands a list of traces to a systematically-named, non-factorised test suite.
func ExpandAll(traces []Trace) testlang.Suite {
	suite := make(testlang.Suite)

	for i, tr := range traces {
		name := fmt.Sprintf("test%d", i)
		suite[name] = tr.Expand(name)
	}

	return suite
}
