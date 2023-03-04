package value_test

import (
	"reflect"
	"testing"

	"github.com/UoY-RoboStar/rtcg/internal/testlang/value"
)

// TestValue_MarshalText tests value text marshaling in several circumstances.
func TestValue_MarshalText(t *testing.T) {
	t.Parallel()

	for name, test := range map[string]struct {
		input value.Value
		want  string
	}{
		"int": {input: value.Int(42), want: "42"},
		"raw": {input: value.Enum("Ok"), want: "Ok"},
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

// TestValue_UnmarshalText tests value text unmarshaling in several circumstances.
func TestValue_UnmarshalText(t *testing.T) {
	t.Parallel()

	for name, test := range map[string]struct {
		input string
		want  value.Value
	}{
		"int": {input: "42", want: value.Int(42)},
		"raw": {input: "Ok", want: value.Enum("Ok")},
		// String trimming tests
		"int-left-pad":  {input: "  42", want: value.Int(42)},
		"int-right-pad": {input: "42  ", want: value.Int(42)},
		"raw-left-pad":  {input: "  Ok", want: value.Enum("Ok")},
		"raw-right-pad": {input: "Ok  ", want: value.Enum("Ok")},
	} {
		input := test.input
		want := test.want

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var got value.Value
			if err := got.UnmarshalText([]byte(input)); err != nil {
				t.Fatalf("unexpected unmarshalling error: %s", err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %s (%v), want %s (%v)", &got, got, &want, want)
			}
		})
	}
}
