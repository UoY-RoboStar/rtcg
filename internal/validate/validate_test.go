package validate_test

import (
	"github.com/UoY-RoboStar/rtcg/internal/testlang"
	"github.com/UoY-RoboStar/rtcg/internal/validate"
	"testing"
)

// TestRoot tests the happy path of Root validation.
func TestRoot(t *testing.T) {
	t.Parallel()

	add := testlang.Output("add", testlang.NoValue())
	sub := testlang.Output("sub", testlang.NoValue())

	branch1 := testlang.Pass(sub, testlang.Fail())
	branch2 := testlang.Pass(add, testlang.Fail())
	branch3 := testlang.Inc(sub, testlang.Inc(add, testlang.Pass(add, testlang.Fail())))

	testlang.MarkAll(&branch1, "test1")
	testlang.MarkAll(&branch2, "test2")
	testlang.MarkAll(&branch3, "test3")

	branch23 := testlang.Inc(add, testlang.Inc(add, branch2, branch3))
	branch23.Mark("test2", "test3")
	branch23.Next[0].Mark("test2", "test3")

	tree := testlang.Root(branch1, branch23)
	tree.Mark("test1", "test2", "test3")

	val, err := validate.Full(&tree)
	if err != nil {
		t.Fatalf("unexpected validation error: %s", err)
	}

	if got := val.Root(); got != &tree {
		t.Fatalf("wrong root node returned: got %v, want %v", got, tree)
	}
}
