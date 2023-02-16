package channel_test

import (
	"reflect"
	"testing"

	"github.com/UoY-RoboStar/rtcg/internal/testlang/channel"
)

// TestKind_MarshalText tests kind marshaling in several circumstances.
func TestKind_MarshalText(t *testing.T) {
	t.Parallel()

	for name, test := range map[string]struct {
		input channel.Kind
		want  string
	}{
		"in":   {input: channel.KindIn, want: "in"},
		"out":  {input: channel.KindOut, want: "out"},
		"call": {input: channel.KindCall, want: "call"},
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
		want  channel.Kind
	}{
		"in":   {input: "in", want: channel.KindIn},
		"out":  {input: "out", want: channel.KindOut},
		"call": {input: "call", want: channel.KindCall},
		"IN":   {input: "IN", want: channel.KindIn},
		"OUT":  {input: "OUT", want: channel.KindOut},
		"CALL": {input: "CALL", want: channel.KindCall},
	} {
		input := test.input
		want := test.want

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var got channel.Kind
			if err := got.UnmarshalText([]byte(input)); err != nil {
				t.Fatalf("unexpected unmarshalling error: %s", err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}
