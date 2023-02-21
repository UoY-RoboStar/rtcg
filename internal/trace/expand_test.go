package trace_test

import (
	"reflect"
	"testing"

	"github.com/UoY-RoboStar/rtcg/internal/testlang"
	"github.com/UoY-RoboStar/rtcg/internal/trace"
)

// TestTrace_Expand tests expansion of traces to trees.
func TestTrace_Expand(t *testing.T) {
	t.Parallel()

	event1 := testlang.Input("foo", testlang.Int(42))
	event2 := testlang.Output("bar", testlang.Raw("baz"))

	for name, test := range map[string]struct {
		input trace.Forbidden
		want  testlang.Node
	}{
		"no-prefix": {
			input: trace.New(event1),
			want:  testlang.TestPoint(event1),
		},
		"lone-prefix": {
			input: trace.New(event1, event2),
			want: testlang.Inc(
				event2,
				testlang.TestPoint(event1)),
		},
	} {
		name := name
		input := test.input
		want := test.want
		testlang.MarkAll(&want, name)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := input.Expand(name)

			if got == nil {
				t.Fatal("got nil")
			}

			if !reflect.DeepEqual(*got, want) {
				t.Fatalf("got %v, want %v", *got, want)
			}
		})
	}
}
