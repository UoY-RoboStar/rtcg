package trace

import (
	"github.com/UoY-RoboStar/rtcg/internal/structure"
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

	n = testlang.Root(n)
	n.Mark(name)

	return &n
}

// ExpandAll expands a suite of Forbidden traces to a non-factorised test suite.
func ExpandAll(traces map[string]Forbidden) testlang.Suite {
	return structure.OverMap(traces, func(k string, t Forbidden) (string, *testlang.Node) {
		return k, t.Expand(k)
	})
}
