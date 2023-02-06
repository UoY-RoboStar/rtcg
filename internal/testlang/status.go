package testlang

import (
	"errors"
	"fmt"
	"strings"
)

// Status is an enumeration of statuses during a test.
type Status uint8

const (
	StatInc  Status = iota // StatInc is the inconclusive status.
	StatFail               // StatFail is the failing status.
	StatPass               // StatPass is the passing status.
)

func (s *Status) MarshalText() (text []byte, err error) {
	switch *s {
	case StatInc:
		text = []byte("inc")
	case StatFail:
		text = []byte("fail")
	case StatPass:
		text = []byte("pass")
	default:
		err = fmt.Errorf("%w: code %d", ErrBadStatus, *s)
	}
	return
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
