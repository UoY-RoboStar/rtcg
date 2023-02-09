package testlang

import (
	"errors"
	"fmt"
	"strings"
)

// Status is an enumeration of statuses during a test.
type Status uint8

const (
	StatusInc  Status = iota // StatusInc is the inconclusive status.
	StatusFail               // StatusFail is the failing status.
	StatusPass               // StatusPass is the passing status.

	FirstStatus = StatusInc  // FirstStatus is the first status.
	LastStatus  = StatusPass // LastStatus is the last status.

	NumStatus = uint8(LastStatus) + 1 // NumStatus is the number of valid status entries.
)

func (s *Status) String() string {
	switch *s {
	case StatusInc:
		return "inc"
	case StatusFail:
		return "fail"
	case StatusPass:
		return "pass"
	default:
		return "unknown"
	}
}

func (s *Status) MarshalText() ([]byte, error) {
	str := s.String()
	if str == "unknown" {
		return nil, fmt.Errorf("%w: code %d", ErrBadStatus, *s)
	}

	return []byte(str), nil
}

func (s *Status) UnmarshalText(text []byte) error {
	textStr := string(text)

	switch {
	case strings.EqualFold(textStr, "inc"):
		*s = StatusInc
	case strings.EqualFold(textStr, "fail"):
		*s = StatusFail
	case strings.EqualFold(textStr, "pass"):
		*s = StatusPass
	default:
		return fmt.Errorf("%w: got %q", ErrBadStatus, textStr)
	}

	return nil
}

// ErrBadStatus occurs when we try to marshal or unmarshal a bad Status value.
var ErrBadStatus = errors.New("unexpected status value")
