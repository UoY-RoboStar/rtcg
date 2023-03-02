// Package strmanip provides string manipulation functions.
package strmanip

import (
	"strings"
	"unicode"
)

// UpcaseFirst upper-cases the first rune in ident.
func UpcaseFirst(ident string) string {
	var build strings.Builder

	// TODO: make this more efficient
	for i, c := range ident {
		if i == 0 {
			_, _ = build.WriteRune(unicode.ToUpper(c))
		} else {
			_, _ = build.WriteRune(c)
		}
	}

	return build.String()
}

// ToLowerUnderscored converts identifier ident to a lowercase_underscored string.
func ToLowerUnderscored(ident string) string {
	var build strings.Builder

	for i, c := range ident {
		if 0 < i && unicode.IsUpper(c) {
			_, _ = build.WriteRune('_')
		}

		_, _ = build.WriteRune(unicode.ToLower(c))
	}

	return build.String()
}
