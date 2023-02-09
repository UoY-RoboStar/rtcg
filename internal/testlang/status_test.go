package testlang_test

import (
	"github.com/UoY-RoboStar/rtcg/internal/testlang"
	"testing"
)

// TestStatus_String tests the stringification of Status.
func TestStatus_String(t *testing.T) {
	for _, test := range []struct {
		input testlang.Status
		want  string
	}{
		{input: testlang.StatInc, want: "inc"},
		{input: testlang.StatFail, want: "fail"},
		{input: testlang.StatPass, want: "pass"},
		{input: testlang.StatPass + 1, want: "unknown"},
	} {
		input := test.input
		want := test.want
		t.Run(want, func(t *testing.T) {
			got := input.String()
			if got != want {
				t.Fatalf("got %q, want %q", got, want)
			}
		})
	}
}
