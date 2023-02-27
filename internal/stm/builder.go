package stm

import (
	"fmt"

	"github.com/UoY-RoboStar/rtcg/internal/structure"
	"github.com/UoY-RoboStar/rtcg/internal/testlang"
	"github.com/UoY-RoboStar/rtcg/internal/validate"
)

// Builder builds state machines from tests.
type Builder struct {
	nodeNum uint64                          // nodeNum is a monotonically increasing counter for naming unnamed nodes.
	stack   structure.Stack[*testlang.Node] // stack is a stack used for in-order test traversal.
	stm     Stm                             // stm is the state machine currently being built.
}

// BuildSuite builds a test suite s into a map of state machines.
func (b *Builder) BuildSuite(s validate.Suite) Suite {
	suite := make(Suite, len(s))

	for k, v := range s {
		m := b.Build(k, v)
		suite[k] = &m
	}

	return suite
}

// Build builds a single state machine from the given validated test.
func (b *Builder) Build(name string, test *validate.Test) Stm {
	b.nodeNum = 0
	b.stm = Stm{States: []*State{}, Tests: structure.NewSet[string]()}

	testRoot := test.Root()
	testRoot.ID = testlang.NodeID(name)

	b.stack.Clear()
	b.stack.Push(testRoot)

	for !b.stack.IsEmpty() {
		node := b.stack.Pop()

		// We don't emit failing states.
		if node.Outcome != testlang.OutcomeFail {
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
	result := NewState(node.ID)

	for i := range node.Next {
		np := &node.Next[i]

		b.ensureNodeID(np)
		b.stm.Tests.Add(np.Tests...)
		result.AddOutgoingNode(np)
	}

	return result
}

func (b *Builder) ensureNodeID(n *testlang.Node) {
	if n.ID == "" {
		n.ID = testlang.NodeID(fmt.Sprintf("node%d", b.nodeNum))
		b.nodeNum++
	}
}
