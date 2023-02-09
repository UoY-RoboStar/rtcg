package testlang_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/UoY-RoboStar/rtcg/internal/testlang"
)

// TestChannel_MarshalText tests event text marshaling in several circumstances.
func TestChannel_MarshalText(t *testing.T) {
	t.Parallel()

	for name, test := range map[string]struct {
		input testlang.Channel
		want  string
	}{
		"in":  {input: testlang.Channel{Name: "foo", Direction: testlang.DirIn}, want: "foo.in"},
		"out": {input: testlang.Channel{Name: "bar", Direction: testlang.DirOut}, want: "bar.out"},
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

// TestChannel_UnmarshalText tests channel text unmarshaling in several circumstances.
func TestChannel_UnmarshalText(t *testing.T) {
	t.Parallel()

	for name, test := range map[string]struct {
		input string
		want  testlang.Channel
		err   error
	}{
		"empty": {input: "", err: testlang.BadEventFieldCountError{Got: 1}},
		"space": {input: "  ", err: testlang.BadEventFieldCountError{Got: 1}},
		"in":    {input: "foo.in", want: testlang.Channel{Name: "foo", Direction: testlang.DirIn}},
		"out":   {input: "bar.out", want: testlang.Channel{Name: "bar", Direction: testlang.DirOut}},
		// String trimming tests
		"left-pad-ch":  {input: "  foo.in", want: testlang.Channel{Name: "foo", Direction: testlang.DirIn}},
		"right-pad-ch": {input: "foo  .in", want: testlang.Channel{Name: "foo", Direction: testlang.DirIn}},
		"left-pad-dir": {
			input: "foo.  in",
			want:  testlang.Channel{Name: "foo", Direction: testlang.DirIn},
		},
		"right-pad-dir": {
			input: "foo.in  ",
			want:  testlang.Channel{Name: "foo", Direction: testlang.DirIn},
		},
	} {
		input := test.input
		want := test.want
		wantErr := test.err

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var got testlang.Channel
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
