// Package strmanip provides string manipulation functions.
package strmanip

import (
	"strings"
	"unicode"
)

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
