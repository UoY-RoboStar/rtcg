package testlang_test

import (
	"reflect"
	"rtcg/internal/testlang"
	"testing"
)

// TestChannel_MarshalText tests event text marshaling in several circumstances.
func TestChannel_MarshalText(t *testing.T) {
	for name, test := range map[string]struct {
		input testlang.Channel
		want  string
	}{
		"empty": {
			input: testlang.Channel{},
			want:  "",
		},
		"in": {
			input: testlang.Channel{Name: "foo", Direction: testlang.In},
			want:  "foo.in",
		},
		"out": {
			input: testlang.Channel{Name: "bar", Direction: testlang.Out},
			want:  "bar.out",
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

// TestChannel_UnmarshalText tests channel text unmarshaling in several circumstances.
func TestChannel_UnmarshalText(t *testing.T) {
	for name, test := range map[string]struct {
		input string
		want  testlang.Channel
	}{
		"empty": {
			input: "",
			want:  testlang.Channel{},
		},
		"space": {
			input: "  ",
			want:  testlang.Channel{},
		},
		"in": {
			input: "foo.in",
			want:  testlang.Channel{Name: "foo", Direction: testlang.In},
		},
		"out": {
			input: "bar.out",
			want:  testlang.Channel{Name: "bar", Direction: testlang.Out},
		},
		// String trimming tests
		"left-pad-ch": {
			input: "  foo.in",
			want:  testlang.Channel{Name: "foo", Direction: testlang.In},
		},
		"right-pad-ch": {
			input: "foo  .in",
			want:  testlang.Channel{Name: "foo", Direction: testlang.In},
		},
		"left-pad-dir": {
			input: "foo.  in",
			want:  testlang.Channel{Name: "foo", Direction: testlang.In},
		},
		"right-pad-dir": {
			input: "foo.in  ",
			want:  testlang.Channel{Name: "foo", Direction: testlang.In},
		},
	} {
		input := test.input
		want := test.want
		t.Run(name, func(t *testing.T) {
			var got testlang.Channel
			if err := got.UnmarshalText([]byte(input)); err != nil {
				t.Fatalf("unexpected unmarshalling error: %s", err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %s (%v), want %s (%v)", &got, got, &want, want)
			}
		})
	}
}
