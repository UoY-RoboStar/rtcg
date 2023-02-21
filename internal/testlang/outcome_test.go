package testlang_test

import (
	"testing"

	"github.com/UoY-RoboStar/rtcg/internal/testlang"
)

// TestStatus_String tests the stringification of Outcome.
func TestStatus_String(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		input testlang.Outcome
		want  string
	}{
		{input: testlang.OutcomeUnset, want: "unset"},
		{input: testlang.OutcomeInc, want: "inc"},
		{input: testlang.OutcomeFail, want: "fail"},
		{input: testlang.OutcomePass, want: "pass"},
		{input: testlang.OutcomePass + 1, want: "Outcome(4)"},
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
