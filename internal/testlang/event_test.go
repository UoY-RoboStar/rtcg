package testlang_test

import (
	"reflect"
	"rtcg/internal/testlang"
	"testing"
)

// TestEvent_MarshalText tests event text marshaling in several circumstances.
func TestEvent_MarshalText(t *testing.T) {
	for name, test := range map[string]struct {
		input testlang.Event
		want  string
	}{
		"empty": {
			input: testlang.Event{},
			want:  "",
		},
		"no-value": {
			input: testlang.Input("foo", testlang.NoValue),
			want:  "foo.in",
		},
		"int-value": {
			input: testlang.Input("foo", testlang.Int(42)),
			want:  "foo.in.42",
		},
		"raw-value": {
			input: testlang.Output("bar", testlang.Raw("Ok")),
			want:  "bar.out.Ok",
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

// TestEvent_UnmarshalText tests event text unmarshaling in several circumstances.
func TestEvent_UnmarshalText(t *testing.T) {
	for name, test := range map[string]struct {
		input string
		want  testlang.Event
	}{
		"empty": {
			input: "",
			want:  testlang.Event{},
		},
		"space": {
			input: "  ",
			want:  testlang.Event{},
		},
		"no-value": {
			input: "foo.in",
			want:  testlang.Input("foo", testlang.NoValue),
		},
		"int-value": {
			input: "foo.in.42",
			want:  testlang.Input("foo", testlang.Int(42)),
		},
		"raw-value": {
			input: "bar.out.Ok",
			want:  testlang.Output("bar", testlang.Raw("Ok")),
		},
		// String trimming tests
		"left-pad-ch": {
			input: "  foo.in.42",
			want:  testlang.Input("foo", testlang.Int(42)),
		},
		"right-pad-ch": {
			input: "foo  .in.42",
			want:  testlang.Input("foo", testlang.Int(42)),
		},
		"left-pad-dir": {
			input: "foo.  in.42",
			want:  testlang.Input("foo", testlang.Int(42)),
		},
		"right-pad-dir": {
			input: "foo.in  .42",
			want:  testlang.Input("foo", testlang.Int(42)),
		},
		"left-pad-val": {
			input: "foo.in.  42",
			want:  testlang.Input("foo", testlang.Int(42)),
		},
		"right-pad-val": {
			input: "foo.in.42  ",
			want:  testlang.Input("foo", testlang.Int(42)),
		},
	} {
		input := test.input
		want := test.want
		t.Run(name, func(t *testing.T) {
			var got testlang.Event
			if err := got.UnmarshalText([]byte(input)); err != nil {
				t.Fatalf("unexpected unmarshalling error: %s", err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %s (%v), want %s (%v)", &got, got, &want, want)
			}
		})
	}
}
