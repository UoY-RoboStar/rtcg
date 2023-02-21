package testlang

import "github.com/UoY-RoboStar/rtcg/internal/structure"

// Walk walks the test tree from node in pre-order, applying function until it returns an error.
func Walk(node *Node, function func(*Node) error) error {
	var stack structure.Stack[*Node]
	stack.Push(node)

	for !stack.IsEmpty() {
		node = stack.Pop()

		if err := function(node); err != nil {
			return err
		}

		// Push in reverse so that the first next-node is serviced first.
		for i := len(node.Next) - 1; 0 <= i; i-- {
			stack.Push(&node.Next[i])
		}
	}

	return nil
}

// MarkAll recursively marks each node in node with the given tests.
func MarkAll(node *Node, tests ...string) {
	_ = Walk(node, func(node *Node) error {
		node.Mark(tests...)

		return nil
	})
}
