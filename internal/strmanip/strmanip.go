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
	var sp formatter
	sp.mode = formatModeLowerUnderscored

	return sp.format(ident)
}

type formatter struct {
	mode      formatMode
	sb        strings.Builder
	lastLower bool
}

func (s *formatter) format(ident string) string {
	for _, char := range ident {
		switch {
		case unicode.IsUpper(char) && s.lastLower:
			s.markSplit()
			s.addChar(char)
		case char == '_':
			s.markSplit()
			s.lastLower = false
		default:
			s.addChar(char)
		}
	}

	return s.sb.String()
}

func (s *formatter) markSplit() {
	if s.mode == formatModeLowerUnderscored {
		s.sb.WriteRune('_')
	}
}

func (s *formatter) addChar(char rune) {
	toWrite := char
	if s.mode == formatModeLowerUnderscored {
		toWrite = unicode.ToLower(toWrite)
	}

	s.lastLower = unicode.IsLower(char)
	_, _ = s.sb.WriteRune(toWrite)
}

type formatMode uint8

const (
	formatModeLowerUnderscored formatMode = iota
)
