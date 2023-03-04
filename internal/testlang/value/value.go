// Package value contains the value model for test traces.
package value

import (
	"bytes"
	"encoding"
	"errors"
	"fmt"
	"reflect"

	"github.com/UoY-RoboStar/rtcg/internal/testlang/rstype"
)

// Value is a value in a trace event.
//
// This type wraps a Valuable to allow it to be safely marshalled and unmarshalled.
// The content of a Value is invariably one of the xValue types provided in this package.
type Value struct {
	content Valuable // content is the content of the value.
}

// None returns an empty value.
func None() Value {
	return New(nil)
}

// New constructs a Value with the given content.
func New(content Valuable) Value {
	return Value{content: content}
}

// IsPresent checks whether v is non-empty.
func (v *Value) IsPresent() bool {
	return !v.IsEmpty()
}

// IsEmpty checks whether v is empty.
func (v *Value) IsEmpty() bool {
	return v.content == nil
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
	return reflect.DeepEqual(v.content, other.content)
}

func (v *Value) MarshalText() ([]byte, error) {
	if v.IsEmpty() {
		return nil, ErrMarshallAbsentValue
	}

	output, err := v.content.MarshalText()
	if err != nil {
		return nil, fmt.Errorf("couldn't marshal internal representation of value: %w", err)
	}

	return output, nil
}

func (v *Value) String() string {
	if v.IsEmpty() {
		return "(empty value)"
	}

	return v.content.String()
}

// StringValue gets a source-interpolatable string representation of this value.
// It differs from String in that it is not meant for human consumption.
func (v *Value) StringValue() string {
	if v.IsEmpty() {
		return ""
	}

	return v.content.StringValue()
}

func (v *Value) UnmarshalText(text []byte) error {
	text = bytes.TrimSpace(text)

	for _, parser := range []func() Valuable{
		func() Valuable {
			tmp := IntValue(0)

			return tryUnmarshal(&tmp, text)
		},
		func() Valuable {
			tmp := EnumValue("")

			return tryUnmarshal(&tmp, text)
		},
	} {
		v.content = parser()
		if v.content != nil {
			return nil
		}
	}

	return fmt.Errorf("%w: input %q", ErrUnsupportedValueType, string(text))
}

func (v *Value) Type() *rstype.RsType {
	if v.IsEmpty() {
		return rstype.Empty()
	}

	return v.content.Type()
}

// Valuable is the type of things that can be treated as content in Value.
type Valuable interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
	fmt.Stringer
	rstype.Typeable

	// StringValue gets a source-interpolatable string representation of this value.
	// It differs from String in that it is not meant for human consumption.
	StringValue() string
}

func tryUnmarshal[V Valuable](tmp V, text []byte) Valuable {
	if err := tmp.UnmarshalText(text); err != nil {
		return nil
	}

	return tmp
}

var (
	ErrMarshallAbsentValue  = errors.New("tried to marshal an absent Value")
	ErrUnsupportedValueType = errors.New("value type not supported")
)
