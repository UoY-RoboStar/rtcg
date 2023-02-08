package stm

import (
	"fmt"
	"rtcg/internal/testlang"
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
func (b *Builder) BuildSuite(s testlang.Suite) map[string]Stm {
	result := make(map[string]Stm, len(s))
	for k, v := range s {
		result[k] = b.Build(k, v)
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

		sn := b.buildNode(n)
		result.Nodes = append(result.Nodes, sn)

		for _, x := range n.Next {
			b.stack.Push(&x)
		}
	}

	return result
}

func (b *Builder) buildNode(n *testlang.Node) Node {
	result := Node{ID: n.ID, Fails: map[string]bool{}}

	for _, x := range n.Next {
		// For failing transitions, we don't emit a destination or value; we just log that this node has failed.
		if x.Status == testlang.StatFail {
			result.IsFail = true
			for _, t := range x.Tests {
				result.Fails[t] = true
			}
			continue
		}

		tr := Transition{Status: x.Status, Value: x.Event.Value, Next: x.ID}
		result.AddTransition(x.Event.Channel, tr)
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
	b.ensureNodeID(n)

	initial := Node{ID: testlang.NodeID(name)}
	tr := Transition{Status: n.Status, Value: n.Event.Value, Next: n.ID}
	initial.AddTransition(n.Event.Channel, tr)

	return Stm{Nodes: []Node{initial}}
}
