package testlang

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

// Direction is an enumeration of communication directions.
type Direction bool

const (
	DirIn  Direction = false // DirIn represents an input.
	DirOut Direction = true  // DirOut represents an output.
)

func (i *Direction) String() string {
	if *i {
		return "out"
	}

	return "in"
}

func (i *Direction) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

func (i *Direction) UnmarshalText(text []byte) error {
	textStr := string(bytes.TrimSpace(text))

	switch {
	case strings.EqualFold(textStr, "in"):
		*i = DirIn
	case strings.EqualFold(textStr, "out"):
		*i = DirOut
	default:
		return fmt.Errorf("%w: got %q", ErrBadInOut, textStr)
	}

	return nil
}

// ErrBadInOut occurs when we try to marshal or unmarshal a bad Direction value.
var ErrBadInOut = errors.New("unexpected in/out value")
