package comm_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/UoY-RoboStar/rtcg/internal/testlang/comm"
)

// TestChannel_MarshalText tests event text marshaling in several circumstances.
func TestChannel_MarshalText(t *testing.T) {
	t.Parallel()

	for name, test := range map[string]struct {
		input comm.Channel
		want  string
	}{
		"in":  {input: comm.Channel{Name: "foo", Kind: comm.KindIn}, want: "foo.in"},
		"out": {input: comm.Channel{Name: "bar", Kind: comm.KindOut}, want: "bar.out"},
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
		want  comm.Channel
		err   error
	}{
		"empty": {input: "", err: comm.BadFieldCountError{Got: 1}},
		"space": {input: "  ", err: comm.BadFieldCountError{Got: 1}},
		"in":    {input: "foo.in", want: comm.Channel{Name: "foo", Kind: comm.KindIn}},
		"out":   {input: "bar.out", want: comm.Channel{Name: "bar", Kind: comm.KindOut}},
		// String trimming tests
		"left-pad-ch":  {input: "  foo.in", want: comm.Channel{Name: "foo", Kind: comm.KindIn}},
		"right-pad-ch": {input: "foo  .in", want: comm.Channel{Name: "foo", Kind: comm.KindIn}},
		"left-pad-dir": {
			input: "foo.  in",
			want:  comm.Channel{Name: "foo", Kind: comm.KindIn},
		},
		"right-pad-dir": {
			input: "foo.in  ",
			want:  comm.Channel{Name: "foo", Kind: comm.KindIn},
		},
	} {
		input := test.input
		want := test.want
		wantErr := test.err

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var got comm.Channel
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
