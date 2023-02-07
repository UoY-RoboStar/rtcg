package testlang

import (
	"encoding"
	"errors"
	"fmt"
	"strconv"
)

// Value is a value in a trace event.
//
// This type wraps a Valuable to allow it to be safely marshalled and unmarshalled.
// The Content of a Value is invariably one of the xValue types provided in this package.
type Value struct {
	Content Valuable // Content is the content of the value.
}

// Valuable is the type of things that can be treated as Content in Value.
type Valuable interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
}

func (v *Value) MarshalText() (text []byte, err error) {
	return v.Content.MarshalText()
}

func (v *Value) UnmarshalText(text []byte) error {
	for _, f := range []func() Valuable{
		func() Valuable {
			tmp := IntValue(0)
			return tryUnmarshal(&tmp, text)
		},
		func() Valuable {
			tmp := RawValue("")
			return tryUnmarshal(&tmp, text)
		},
	} {
		v.Content = f()
		if v.Content != nil {
			return nil
		}
	}
	return fmt.Errorf("%w: input %q", ErrUnsupportedValueType, string(text))
}

func tryUnmarshal[V Valuable](tmp V, text []byte) Valuable {
	if err := tmp.UnmarshalText(text); err != nil {
		return nil
	}
	return tmp
}

// IntValue is a value that is an integer.
type IntValue int64

func (i *IntValue) MarshalText() (text []byte, err error) {
	return strconv.AppendInt(text, int64(*i), 10), nil
}

func (i *IntValue) UnmarshalText(text []byte) error {
	var err error
	*((*int64)(i)), err = strconv.ParseInt(string(text), 10, 64)
	return err
}

// RawValue is a value that is an uninterpreted string.
// This usually suggests that the parser has given up trying to parse it as something else.
type RawValue string

func (r *RawValue) MarshalText() (text []byte, err error) {
	// TODO: escape anything that could make this a non-raw value?
	return []byte(*r), nil
}

func (r *RawValue) UnmarshalText(text []byte) error {
	// TODO: refuse anything that could make this a non-raw value?
	*r = RawValue(text)
	return nil
}

var ErrUnsupportedValueType = errors.New("value type not supported")
