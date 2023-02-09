package testlang

import (
	"errors"
	"fmt"
	"strings"
)

// Status is an enumeration of statuses during a test.
type Status uint8

const (
	StatInc   Status = iota // StatInc is the inconclusive status.
	StatFail                // StatFail is the failing status.
	StatPass                // StatPass is the passing status.
	NumStatus               // NumStatus is the number of valid status entries.
)

// AllStatuses contains every possible test status in order.
var AllStatuses = [NumStatus]Status{StatInc, StatFail, StatPass}

func (s *Status) String() string {
	switch *s {
	case StatInc:
		return "inc"
	case StatFail:
		return "fail"
	case StatPass:
		return "pass"
	default:
		return "unknown"
	}
}

func (s *Status) MarshalText() (text []byte, err error) {
	str := s.String()
	if str == "unknown" {
		return nil, fmt.Errorf("%w: code %d", ErrBadStatus, *s)
	}
	return []byte(str), nil
}

func (s *Status) UnmarshalText(text []byte) error {
	textStr := string(text)
	if strings.EqualFold(textStr, "inc") {
		*s = StatInc
	} else if strings.EqualFold(textStr, "fail") {
		*s = StatFail
	} else if strings.EqualFold(textStr, "pass") {
		*s = StatPass
	} else {
		return fmt.Errorf("%w: got %q", ErrBadStatus, textStr)
	}
	return nil
}

// ErrBadStatus occurs when we try to marshal or unmarshal a bad Status value.
var ErrBadStatus = errors.New("unexpected status value")
