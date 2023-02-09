package stm

import (
	"fmt"
	"github.com/UoY-RoboStar/rtcg/internal/testlang"
)

// Builder builds state machines from tests.
type Builder struct {
	nodeNum uint64                // nodeNum is a monotonically increasing counter for naming unnamed nodes.
	stack   Stack[*testlang.Node] // stack is a stack used for in-order test traversal.
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
	result := make(Suite, len(s))
	for k, v := range s {
		m := b.Build(k, v)
		result[k] = &m
	}
	return result
}

// Build builds a single test from testRoot onwards.
func (b *Builder) Build(name string, testRoot *testlang.Node) Stm {
	// We don't reset nodeNum, in case we're building a whole suite.

	result := b.initStm(name, testRoot)

	b.stack.Clear()
	b.stack.Push(testRoot)

	for !b.stack.IsEmpty() {
		n := b.stack.Pop()
		b.ensureNodeIDs(n)

		// We don't emit failing states.
		if n.Status == testlang.StatFail {
			continue
		}

		sn := b.buildState(n)
		result.States = append(result.States, sn)

		for _, x := range n.Next {
			b.stack.Push(&x)
		}
	}

	return result
}

func (b *Builder) buildState(n *testlang.Node) *State {
	result := NewState(n.ID)

	for _, x := range n.Next {
		// For failing transitions, we don't emit a destination or value; we just log that this node has failed.
		result.AddVerdictsFromNode(x)
		if !result.Fail().IsObserved {
			result.AddTransitionToNode(&x)
		}
	}

	return result
}

// ensureNodeIDs makes sure that n and all of its immediate descendants have NodeIDs defined.
func (b *Builder) ensureNodeIDs(n *testlang.Node) {
	b.ensureNodeID(n)
	for i := range n.Next {
		b.ensureNodeID(&n.Next[i])
	}
}

func (b *Builder) ensureNodeID(n *testlang.Node) {
	if n.ID == "" {
		n.ID = testlang.NodeID(fmt.Sprintf("node_%d", b.nodeNum))
		b.nodeNum++
	}
}

func (b *Builder) initStm(name string, n *testlang.Node) Stm {
	// Ideally, we wouldn't do this, and, instead, we'd just create a special prefix node and put that in the stack.
	// However, that would require us to copy n into its Next list.
	b.ensureNodeID(n)

	initial := NewState(testlang.NodeID(name))
	initial.AddTransitionToNode(n)

	return Stm{States: []*State{initial}}
}
