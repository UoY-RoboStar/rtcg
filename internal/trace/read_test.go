package trace_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/UoY-RoboStar/rtcg/internal/testlang"
	"github.com/UoY-RoboStar/rtcg/internal/testlang/value"
	"github.com/UoY-RoboStar/rtcg/internal/trace"
)

const readTestCase = `
bar.out
foo.in, bar.out
named: foo.in, baz.out.2, foo.in, bar.out
`

// TestRead tests the happy path of reading a trace.
func TestRead(t *testing.T) {
	t.Parallel()

	// TODO: more test cases?

	input := strings.NewReader(readTestCase)

	got, err := trace.Read(input)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	barOut := testlang.Output("bar", value.None())
	fooIn := testlang.Input("foo", value.None())
	bazOut := testlang.Output("baz", value.Int(2))

	want := []trace.Forbidden{
		trace.New().Forbid(barOut),
		trace.New(fooIn).Forbid(barOut),
		trace.New(fooIn, bazOut, fooIn).ForbidWithName(barOut, "named"),
	}

	if len(got) != len(want) {
		t.Fatalf("got %d traces, want %d", len(got), len(want))
	}

	for i, g := range got {
		if !reflect.DeepEqual(g, want[i]) {
			t.Errorf("traces at index %d differ: got %v, want %v", i, g, want[i])
		}
	}
}
