package stm

import (
	"fmt"

	"github.com/UoY-RoboStar/rtcg/internal/structure"
	"github.com/UoY-RoboStar/rtcg/internal/testlang"
)

// Builder builds state machines from tests.
type Builder struct {
	nodeNum uint64                          // nodeNum is a monotonically increasing counter for naming unnamed nodes.
	stack   structure.Stack[*testlang.Node] // stack is a stack used for in-order test traversal.
	stm     Stm                             // stm is the state machine currently being built.
}

// BuildSuite builds a test suite s into a map of state machines.
func (b *Builder) BuildSuite(s testlang.Suite) Suite {
	suite := make(Suite, len(s))

	for k, v := range s {
		m := b.Build(k, v)
		suite[k] = &m
	}

	return suite
}

// Build builds a single test from testRoot onwards.
func (b *Builder) Build(name string, testRoot *testlang.Node) Stm {
	// We don't reset nodeNum, in case we're building a whole suite.
	b.initStm(name, testRoot)

	b.stack.Clear()
	b.stack.Push(testRoot)

	for !b.stack.IsEmpty() {
		node := b.stack.Pop()

		// We don't emit failing states.
		if node.Status != testlang.StatusFail {
			b.processNode(node)
		}
	}

	return b.stm
}

func (b *Builder) processNode(node *testlang.Node) {
	sn := b.buildState(node)
	b.stm.States = append(b.stm.States, sn)

	for i := range node.Next {
		b.stack.Push(&node.Next[i])
	}
}

func (b *Builder) buildState(node *testlang.Node) *State {
	b.ensureNodeID(node)
	result := NewState(node.ID)

	for i := range node.Next {
		np := &node.Next[i]

		b.stm.Tests.Add(np.Tests...)
		result.AddOutgoingNode(np)
	}

	return result
}

func (b *Builder) ensureNodeID(n *testlang.Node) {
	if n.ID == "" {
		n.ID = testlang.NodeID(fmt.Sprintf("node_%d", b.nodeNum))
		b.nodeNum++
	}
}

func (b *Builder) initStm(name string, node *testlang.Node) {
	// Ideally, we wouldn't do this, and, instead, we'd just create a special prefix node and put that in the stack.
	// However, that would require us to copy node into its Next list.
	b.ensureNodeID(node)

	initial := NewState(testlang.NodeID(name))
	initial.AddTransitionToNode(node)

	b.stm = Stm{States: []*State{initial}, Tests: structure.Set[string]{}}
	b.stm.Tests.Add(node.Tests...)
}
