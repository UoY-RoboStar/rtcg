// Package validate contains validators for rtcg tests.
package validate

import (
	"fmt"

	"github.com/UoY-RoboStar/rtcg/internal/testlang"
)

// Test represents a validated test.
type Test struct {
	root *testlang.Node
}

// Root gets a pointer to the root node of a validated Test.
func (t Test) Root() *testlang.Node {
	return t.root
}

// Full performs a full validation of the test starting at root.
// If successful, it returns a Test.
func Full(root *testlang.Node) (*Test, error) {
	// TODO: check test monotonicity and partitioning
	err := testlang.Walk(root, func(node *testlang.Node) error {
		var err error

		if node == root {
			err = Root(node)
		} else {
			err = Node(node)
		}

		if err != nil {
			return fmt.Errorf("invalid node %q: %w", node, err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to validate test: %w", err)
	}

	return &Test{root: root}, nil
}

// OneNext tries to assert that there is exactly one next node in node, and gets it if successful.
//
// Passing nodes need to have exactly one next node (the failing node), but we export this
// function as it is useful for collapsing trace tests.
func OneNext(node *testlang.Node) (*testlang.Node, error) {
	if nnodes := len(node.Next); nnodes != 1 {
		return nil, fmt.Errorf("%w: node count was %d", ErrNeedOneNode, nnodes)
	}

	return &node.Next[0], nil
}
