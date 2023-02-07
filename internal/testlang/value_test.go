package testlang_test

import (
	"reflect"
	"rtcg/internal/testlang"
	"testing"
)

// TestValue_MarshalText tests value text marshaling in several circumstances.
func TestValue_MarshalText(t *testing.T) {
	for name, test := range map[string]struct {
		input testlang.Value
		want  string
	}{
		"int": {
			input: testlang.Int(42),
			want:  "42",
		},
		"raw": {
			input: testlang.Raw("Ok"),
			want:  "Ok",
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

// TestValue_UnmarshalText tests value text unmarshaling in several circumstances.
func TestValue_UnmarshalText(t *testing.T) {
	for name, test := range map[string]struct {
		input string
		want  testlang.Value
	}{
		"int": {
			input: "42",
			want:  testlang.Int(42),
		},
		"raw": {
			input: "Ok",
			want:  testlang.Raw("Ok"),
		},
		// String trimming tests
		"int-left-pad": {
			input: "  42",
			want:  testlang.Int(42),
		},
		"int-right-pad": {
			input: "42  ",
			want:  testlang.Int(42),
		},
		"raw-left-pad": {
			input: "  Ok",
			want:  testlang.Raw("Ok"),
		},
		"raw-right-pad": {
			input: "Ok  ",
			want:  testlang.Raw("Ok"),
		},
	} {
		input := test.input
		want := test.want
		t.Run(name, func(t *testing.T) {
			var got testlang.Value
			if err := got.UnmarshalText([]byte(input)); err != nil {
				t.Fatalf("unexpected unmarshalling error: %s", err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}
