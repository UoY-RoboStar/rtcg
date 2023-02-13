package trace

import (
	"fmt"
	"github.com/UoY-RoboStar/rtcg/internal/testlang"
)

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
