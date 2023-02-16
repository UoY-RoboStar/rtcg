package channel_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/UoY-RoboStar/rtcg/internal/testlang/channel"
)

// TestChannel_MarshalText tests event text marshaling in several circumstances.
func TestChannel_MarshalText(t *testing.T) {
	t.Parallel()

	for name, test := range map[string]struct {
		input channel.Channel
		want  string
	}{
		"in":  {input: channel.Channel{Name: "foo", Kind: channel.KindIn}, want: "foo.in"},
		"out": {input: channel.Channel{Name: "bar", Kind: channel.KindOut}, want: "bar.out"},
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
		want  channel.Channel
		err   error
	}{
		"empty": {input: "", err: channel.BadFieldCountError{Got: 1}},
		"space": {input: "  ", err: channel.BadFieldCountError{Got: 1}},
		"in":    {input: "foo.in", want: channel.Channel{Name: "foo", Kind: channel.KindIn}},
		"out":   {input: "bar.out", want: channel.Channel{Name: "bar", Kind: channel.KindOut}},
		// String trimming tests
		"left-pad-ch":  {input: "  foo.in", want: channel.Channel{Name: "foo", Kind: channel.KindIn}},
		"right-pad-ch": {input: "foo  .in", want: channel.Channel{Name: "foo", Kind: channel.KindIn}},
		"left-pad-dir": {
			input: "foo.  in",
			want:  channel.Channel{Name: "foo", Kind: channel.KindIn},
		},
		"right-pad-dir": {
			input: "foo.in  ",
			want:  channel.Channel{Name: "foo", Kind: channel.KindIn},
		},
	} {
		input := test.input
		want := test.want
		wantErr := test.err

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var got channel.Channel
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
