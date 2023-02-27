package trace

import (
	"github.com/UoY-RoboStar/rtcg/internal/structure"
	"github.com/UoY-RoboStar/rtcg/internal/testlang"
)

// Expand expands a single Forbidden trace into a test, tagged throughout with name.
func (t Forbidden) Expand(name string) *testlang.Node {
	// Work backwards through the trace, building the tree from the failure.
	node := testlang.TestPoint(t.Forbid, name)

	for i := len(t.Prefix) - 1; 0 <= i; i-- {
		node = testlang.Inc(t.Prefix[i], node)
		node.Mark(name)
	}

	node = testlang.Root(node)
	node.Mark(name)

	return &node
}

// ExpandAll expands a suite of Forbidden traces to a non-factorised test suite.
func ExpandAll(traces map[string]Forbidden) testlang.Suite {
	return structure.OverMap(traces, func(k string, t Forbidden) (string, *testlang.Node) {
		return k, t.Expand(k)
	})
}
