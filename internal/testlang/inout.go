package testlang

import (
	"bytes"
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

func (i *InOut) String() string {
	if *i {
		return "out"
	}
	return "in"
}

func (i *InOut) MarshalText() (text []byte, err error) {
	return []byte(i.String()), nil
}

func (i *InOut) UnmarshalText(text []byte) error {
	textStr := string(bytes.TrimSpace(text))
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
