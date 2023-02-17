package trace_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/UoY-RoboStar/rtcg/internal/testlang"
	"github.com/UoY-RoboStar/rtcg/internal/trace"
)

const readTestCase = `
bar.out
foo.in, bar.out
foo.in, baz.out.2, foo.in, bar.out
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

	barOut := testlang.Output("bar", testlang.NoValue())
	fooIn := testlang.Input("foo", testlang.NoValue())
	bazOut := testlang.Output("baz", testlang.Int(2))

	want := []trace.Forbidden{
		trace.New(barOut),
		trace.New(barOut, fooIn),
		trace.New(barOut, fooIn, bazOut, fooIn),
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
