package testlang

import (
	"bytes"
	"encoding"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// Value is a value in a trace event.
//
// This type wraps a Valuable to allow it to be safely marshalled and unmarshalled.
// The Content of a Value is invariably one of the xValue types provided in this package.
type Value struct {
	Content Valuable // Content is the content of the value.
}

// NoValue returns an empty value.
func NoValue() Value {
	return Value{Content: nil}
}

// IsPresent checks whether v is non-empty.
func (v *Value) IsPresent() bool {
	return !v.IsEmpty()
}

// IsEmpty checks whether v is empty.
func (v *Value) IsEmpty() bool {
	return v.Content == nil
}

// Equals tests equality on two Values.
func (v *Value) Equals(other Value) bool {
	if v.IsEmpty() {
		return other.IsEmpty()
	}

	if other.IsEmpty() {
		return false
	}

	// TODO: is this the right way to compare these?
	return reflect.DeepEqual(v.Content, other.Content)
}

func (v *Value) MarshalText() ([]byte, error) {
	if v.IsEmpty() {
		return nil, ErrMarshallAbsentValue
	}

	output, err := v.Content.MarshalText()
	if err != nil {
		return nil, fmt.Errorf("couldn't marshal internal representation of value: %w", err)
	}

	return output, nil
}

func (v *Value) String() string {
	if v.IsEmpty() {
		return "(empty value)"
	}

	return v.Content.String()
}

func (v *Value) UnmarshalText(text []byte) error {
	text = bytes.TrimSpace(text)

	for _, parser := range []func() Valuable{
		func() Valuable {
			tmp := IntValue(0)

			return tryUnmarshal(&tmp, text)
		},
		func() Valuable {
			tmp := RawValue("")

			return tryUnmarshal(&tmp, text)
		},
	} {
		v.Content = parser()
		if v.Content != nil {
			return nil
		}
	}

	return fmt.Errorf("%w: input %q", ErrUnsupportedValueType, string(text))
}

// Valuable is the type of things that can be treated as Content in Value.
type Valuable interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
	fmt.Stringer
}

func tryUnmarshal[V Valuable](tmp V, text []byte) Valuable {
	if err := tmp.UnmarshalText(text); err != nil {
		return nil
	}

	return tmp
}

// IntValue is a value that is an integer.
type IntValue int64

// intValueBase encodes the base used for parsing and emitting integer values.
const intValueBase = 10

func (i *IntValue) MarshalText() ([]byte, error) {
	return strconv.AppendInt([]byte{}, int64(*i), intValueBase), nil
}

func (i *IntValue) UnmarshalText(text []byte) error {
	var err error
	*((*int64)(i)), err = strconv.ParseInt(string(text), intValueBase, 64)

	if err != nil {
		return fmt.Errorf("couldn't parse %q as int: %w", string(text), err)
	}

	return nil
}

func (i *IntValue) String() string {
	return fmt.Sprintf("int!%d", int64(*i))
}

// Int constructs an integer Value.
func Int(i int64) Value {
	c := IntValue(i)

	return Value{Content: &c}
}

// RawValue is a value that is an uninterpreted string.
// This usually suggests that the parser has given up trying to parse it as something else.
type RawValue string

func (r *RawValue) MarshalText() ([]byte, error) {
	// TODO: escape anything that could make this a non-raw value?
	return []byte(*r), nil
}

func (r *RawValue) UnmarshalText(text []byte) error {
	// TODO: refuse anything that could make this a non-raw value?
	*r = RawValue(text)

	return nil
}

func (r *RawValue) String() string {
	return fmt.Sprintf("raw!%q", string(*r))
}

// Raw constructs an uninterpreted Value.
func Raw(contents string) Value {
	c := RawValue(contents)

	return Value{Content: &c}
}

var (
	ErrMarshallAbsentValue  = errors.New("tried to marshal an absent Value")
	ErrUnsupportedValueType = errors.New("value type not supported")
)
