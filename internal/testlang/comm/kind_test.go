package comm_test

import (
	"reflect"
	"testing"

	"github.com/UoY-RoboStar/rtcg/internal/testlang/comm"
)

// TestKind_MarshalText tests kind marshaling in several circumstances.
func TestKind_MarshalText(t *testing.T) {
	t.Parallel()

	for name, test := range map[string]struct {
		input comm.Kind
		want  string
	}{
		"in":   {input: comm.KindIn, want: "in"},
		"out":  {input: comm.KindOut, want: "out"},
		"call": {input: comm.KindCall, want: "call"},
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

// TestKind_UnmarshalText tests in/out text unmarshaling in several circumstances.
func TestKind_UnmarshalText(t *testing.T) {
	t.Parallel()

	for name, test := range map[string]struct {
		input string
		want  comm.Kind
	}{
		"in":   {input: "in", want: comm.KindIn},
		"out":  {input: "out", want: comm.KindOut},
		"call": {input: "call", want: comm.KindCall},
		"IN":   {input: "IN", want: comm.KindIn},
		"OUT":  {input: "OUT", want: comm.KindOut},
		"CALL": {input: "CALL", want: comm.KindCall},
	} {
		input := test.input
		want := test.want

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var got comm.Kind
			if err := got.UnmarshalText([]byte(input)); err != nil {
				t.Fatalf("unexpected unmarshalling error: %s", err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}
