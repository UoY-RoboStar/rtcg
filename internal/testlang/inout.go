package testlang

import (
	"errors"
	"fmt"
	"strings"
)

// InOut is an enumeration of communication directions.
type InOut bool

const (
	In  InOut = false // In represents an input.
	Out InOut = true  // Out represents output.
)

func (i *InOut) MarshalText() (text []byte, err error) {
	switch *i {
	case In:
		text = []byte("in")
	case Out:
		text = []byte("out")
	}
	// *i is a bool, so the above should be the only two allowed values
	return
}

func (i *InOut) UnmarshalText(text []byte) error {
	textStr := string(text)
	if strings.EqualFold(textStr, "in") {
		*i = In
	} else if strings.EqualFold(textStr, "out") {
		*i = Out
	} else {
		return fmt.Errorf("%w: got %q", ErrBadInOut, textStr)
	}
	return nil
}

// ErrBadInOut occurs when we try to marshal or unmarshal a bad InOut value.
var ErrBadInOut = errors.New("unexpected in/out value")
