package testlang_test

import (
	"reflect"
	"testing"

	"github.com/UoY-RoboStar/rtcg/internal/testlang"
)

// TestInOut_MarshalText tests in/out marshaling in several circumstances.
func TestInOut_MarshalText(t *testing.T) {
	t.Parallel()

	for name, test := range map[string]struct {
		input testlang.Direction
		want  string
	}{
		"in":  {input: testlang.DirIn, want: "in"},
		"out": {input: testlang.DirOut, want: "out"},
	} {
		input := test.input
		want := test.want

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			bs, err := input.MarshalText()
			if err != nil {
				t.Fatalf("unexpected marshalling error: %s", err)
			}
			got := string(bs)
			if got != want {
				t.Fatalf("got %q, want %q", got, want)
			}
		})
	}
}

// TestInOut_UnmarshalText tests in/out text unmarshaling in several circumstances.
func TestInOut_UnmarshalText(t *testing.T) {
	t.Parallel()

	for name, test := range map[string]struct {
		input string
		want  testlang.Direction
	}{
		"in":  {input: "in", want: testlang.DirIn},
		"out": {input: "out", want: testlang.DirOut},
		"IN":  {input: "IN", want: testlang.DirIn},
		"OUT": {input: "OUT", want: testlang.DirOut},
		// String trimming tests
		"in-left-pad":   {input: "  in", want: testlang.DirIn},
		"in-right-pad":  {input: "in  ", want: testlang.DirIn},
		"out-left-pad":  {input: "  out", want: testlang.DirOut},
		"out-right-pad": {input: "out  ", want: testlang.DirOut},
	} {
		input := test.input
		want := test.want

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var got testlang.Direction
			if err := got.UnmarshalText([]byte(input)); err != nil {
				t.Fatalf("unexpected unmarshalling error: %s", err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}
