package value

import (
	"fmt"
	"strconv"

	"github.com/UoY-RoboStar/rtcg/internal/testlang/rstype"
)

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

func (i *IntValue) StringValue() string {
	return strconv.FormatInt(int64(*i), intValueBase)
}

func (i *IntValue) Type() *rstype.RsType {
	if 0 <= *i {
		return rstype.Arithmos(rstype.ArithmosDomainNat)
	}

	return rstype.Arithmos(rstype.ArithmosDomainInt)
}

// Int constructs an integer Value.
func Int(i int64) Value {
	iv := IntValue(i)

	return New(&iv)
}
