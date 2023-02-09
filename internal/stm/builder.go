package stm

import (
	"fmt"

	"github.com/UoY-RoboStar/rtcg/internal/testlang"
)

// Builder builds state machines from tests.
type Builder struct {
	nodeNum uint64                // nodeNum is a monotonically increasing counter for naming unnamed nodes.
	stack   Stack[*testlang.Node] // stack is a stack used for in-order test traversal.
	current Stm                   // current is the state machine currently being built.
}

// Stack is a(n inefficient) implementation of a stack.
type Stack[T any] struct {
	items []T
	count int
}

// Push pushes item onto the stack.
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
	s.count++
}

// Pop pops an item from s.
//
// It does not check to see if the stack is empty.
func (s *Stack[T]) Pop() T {
	item := s.items[s.count-1]
	s.items = s.items[:s.count-1]
	s.count--

	return item
}

// IsEmpty is true iff the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return s.count == 0
}

// Clear empties stack s.
func (s *Stack[T]) Clear() {
	s.items = []T{}
	s.count = 0
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

	return b.current
}

func (b *Builder) processNode(node *testlang.Node) {
	sn := b.buildState(node)
	b.current.States = append(b.current.States, sn)

	for i := range node.Next {
		b.stack.Push(&node.Next[i])
	}
}

func (b *Builder) buildState(node *testlang.Node) *State {
	b.ensureNodeID(node)
	result := NewState(node.ID)

	for i, n := range node.Next {
		// For failing transitions, we don't emit a destination or value; we just log that this node has failed.
		result.AddVerdictsFromNode(n)

		if !result.Fail().IsObserved {
			result.AddTransitionToNode(&node.Next[i])
		}
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

	b.current = Stm{States: []*State{initial}}
}
