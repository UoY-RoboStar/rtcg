package testlang_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/UoY-RoboStar/rtcg/internal/testlang"
	"github.com/UoY-RoboStar/rtcg/internal/testlang/channel"
	"github.com/UoY-RoboStar/rtcg/internal/testlang/value"
)

// TestEvent_MarshalText tests event text marshaling in several circumstances.
func TestEvent_MarshalText(t *testing.T) {
	t.Parallel()

	for name, test := range map[string]struct {
		input testlang.Event
		want  string
	}{
		"no-value": {
			input: testlang.Input("foo", value.None()),
			want:  "foo.in",
		},
		"int-value": {
			input: testlang.Input("foo", value.Int(42)),
			want:  "foo.in.42",
		},
		"raw-value": {
			input: testlang.Output("bar", value.Enum("Ok")),
			want:  "bar.out.Ok",
		},
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

// TestEvent_UnmarshalText tests event text unmarshaling in several circumstances.
func TestEvent_UnmarshalText(t *testing.T) {
	t.Parallel()

	for name, test := range map[string]struct {
		input string
		want  testlang.Event
		err   error
	}{
		"empty":     {input: "", err: channel.BadFieldCountError{Got: 1}},
		"space":     {input: "  ", err: channel.BadFieldCountError{Got: 1}},
		"no-value":  {input: "foo.in", want: testlang.Input("foo", value.None())},
		"int-value": {input: "foo.in.42", want: testlang.Input("foo", value.Int(42))},
		"raw-value": {input: "bar.out.Ok", want: testlang.Output("bar", value.Enum("Ok"))},
		// String trimming tests
		"left-pad-ch":   {input: "  foo.in.42", want: testlang.Input("foo", value.Int(42))},
		"right-pad-ch":  {input: "foo  .in.42", want: testlang.Input("foo", value.Int(42))},
		"left-pad-dir":  {input: "foo.  in.42", want: testlang.Input("foo", value.Int(42))},
		"right-pad-dir": {input: "foo.in  .42", want: testlang.Input("foo", value.Int(42))},
		"left-pad-val":  {input: "foo.in.  42", want: testlang.Input("foo", value.Int(42))},
		"right-pad-val": {input: "foo.in.42  ", want: testlang.Input("foo", value.Int(42))},
	} {
		input := test.input
		want := test.want
		wantErr := test.err

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var got testlang.Event
			err := got.UnmarshalText([]byte(input))
			if err != nil && !errors.Is(err, wantErr) {
				t.Fatalf("unexpected unmarshalling error: %s", err)
			}
			if err == nil && wantErr != nil {
				t.Fatalf("expected unmarshalling error %q, but got none", wantErr)
			}
			if wantErr == nil && !reflect.DeepEqual(got, want) {
				t.Fatalf("got %s (%v), want %s (%v)", &got, got, &want, want)
			}
		})
	}
}
