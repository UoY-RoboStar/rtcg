package value

import (
	"fmt"
	"github.com/UoY-RoboStar/rtcg/internal/testlang/rstype"
)

// EnumValue is a value that is an enumeration constant.
// This usually suggests that the parser has given up trying to parse it as something else.
type EnumValue string

func (e *EnumValue) MarshalText() ([]byte, error) {
	// TODO: escape anything that could make this a non-raw value?
	return []byte(*e), nil
}

func (e *EnumValue) UnmarshalText(text []byte) error {
	// TODO: refuse anything that could make this a non-raw value?
	*e = EnumValue(text)

	return nil
}

func (e *EnumValue) String() string {
	return fmt.Sprintf("enum!%q", string(*e))
}

func (e *EnumValue) StringValue() string {
	return string(*e)
}

func (*EnumValue) Type() *rstype.RsType {
	return rstype.Enum()
}

// Enum constructs an enumeration constant Value.
func Enum(contents string) Value {
	enum := EnumValue(contents)

	return New(&enum)
}
