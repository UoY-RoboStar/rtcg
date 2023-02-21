package trace

import (
	"fmt"

	"github.com/UoY-RoboStar/rtcg/internal/testlang"
)

// Expand expands a single Forbidden trace into a test, tagged throughout with name.
func (t Forbidden) Expand(name string) *testlang.Node {
	// Work backwards through the trace, building the tree from the failure.
	n := testlang.TestPoint(t.Forbid, name)

	for i := len(t.Prefix) - 1; 0 <= i; i-- {
		n = testlang.Inc(t.Prefix[i], n)
		n.Mark(name)
	}

	return &n
}

// ExpandAll expands a list of Forbidden traces to a systematically-named, non-factorised test suite.
func ExpandAll(traces []Forbidden) testlang.Suite {
	suite := make(testlang.Suite)

	for i, tr := range traces {
		name := fmt.Sprintf("test%d", i)
		suite[name] = tr.Expand(name)
	}

	return suite
}
