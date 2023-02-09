package testlang_test

import (
	"github.com/UoY-RoboStar/rtcg/internal/testlang"
	"reflect"
	"testing"
)

// TestInOut_MarshalText tests in/out marshaling in several circumstances.
func TestInOut_MarshalText(t *testing.T) {
	for name, test := range map[string]struct {
		input testlang.InOut
		want  string
	}{
		"in": {
			input: testlang.In,
			want:  "in",
		},
		"out": {
			input: testlang.Out,
			want:  "out",
		},
	} {
		input := test.input
		want := test.want
		t.Run(name, func(t *testing.T) {
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
	for name, test := range map[string]struct {
		input string
		want  testlang.InOut
	}{
		"in":  {input: "in", want: testlang.In},
		"out": {input: "out", want: testlang.Out},
		"IN":  {input: "IN", want: testlang.In},
		"OUT": {input: "OUT", want: testlang.Out},
		// String trimming tests
		"in-left-pad":   {input: "  in", want: testlang.In},
		"in-right-pad":  {input: "in  ", want: testlang.In},
		"out-left-pad":  {input: "  out", want: testlang.Out},
		"out-right-pad": {input: "out  ", want: testlang.Out},
	} {
		input := test.input
		want := test.want
		t.Run(name, func(t *testing.T) {
			var got testlang.InOut
			if err := got.UnmarshalText([]byte(input)); err != nil {
				t.Fatalf("unexpected unmarshalling error: %s", err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}
