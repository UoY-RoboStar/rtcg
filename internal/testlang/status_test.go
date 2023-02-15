package testlang_test

import (
	"testing"

	"github.com/UoY-RoboStar/rtcg/internal/testlang"
)

// TestStatus_String tests the stringification of Status.
func TestStatus_String(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		input testlang.Status
		want  string
	}{
		{input: testlang.StatusInc, want: "inc"},
		{input: testlang.StatusFail, want: "fail"},
		{input: testlang.StatusPass, want: "pass"},
		{input: testlang.StatusPass + 1, want: "Status(3)"},
	} {
		input := test.input
		want := test.want
		t.Run(want, func(t *testing.T) {
			t.Parallel()

			got := input.String()
			if got != want {
				t.Fatalf("got %q, want %q", got, want)
			}
		})
	}
}
