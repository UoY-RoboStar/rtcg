package strmanip_test

import (
	"testing"

	"github.com/UoY-RoboStar/rtcg/internal/strmanip"
)

// TestToLowerUnderscored tests ToLowerUnderscored.
func TestToLowerUnderscored(t *testing.T) {
	t.Parallel()

	tests := map[string]string{
		"lower_underscored": "lower_underscored",
		"UPPER_UNDERSCORED": "upper_underscored",
		"Mixed_Underscored": "mixed_underscored",
		"camelCase":         "camel_case",
		"PascalCase":        "pascal_case",
	}
	for input, want := range tests {
		input := input
		want := want

		t.Run(input, func(t *testing.T) {
			t.Parallel()

			if got := strmanip.ToLowerUnderscored(input); got != want {
				t.Errorf("ToLowerUnderscored(%q) = %q, want %q", input, got, want)
			}
		})
	}
}

// TestUpcaseFirst tests UpcaseFirst.
func TestUpcaseFirst(t *testing.T) {
	t.Parallel()

	tests := map[string]string{
		"lower_underscored": "Lower_underscored",
		"UPPER_UNDERSCORED": "UPPER_UNDERSCORED",
		"Mixed_Underscored": "Mixed_Underscored",
		"camelCase":         "CamelCase",
		"PascalCase":        "PascalCase",
	}
	for input, want := range tests {
		input := input
		want := want

		t.Run(input, func(t *testing.T) {
			t.Parallel()

			if got := strmanip.UpcaseFirst(input); got != want {
				t.Errorf("UpcaseFirst(%q) = %q, want %q", input, got, want)
			}
		})
	}
}
