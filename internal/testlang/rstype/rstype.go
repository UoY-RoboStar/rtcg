// Package rstype encodes a part of the RoboStar (Z) type system.
package rstype

import (
	"bytes"
	"fmt"
)

// RsType defines a RoboStar type.
type RsType struct {
	kind           Kind           // Kind gets the underlying Kind of this type.
	arithmosDomain ArithmosDomain // ArithmosDomain gets the underlying domain, if Kind is Arithmos.
}

// IsEmpty gets whether this type is empty.
func (t *RsType) IsEmpty() bool {
	return t.kind == KindEmpty
}

// IsEnum gets whether this type is an enum.
func (t *RsType) IsEnum() bool {
	return t.kind == KindEnum
}

// IsNat gets whether this type is that of natural numbers.
func (t *RsType) IsNat() bool {
	return t.IsArithmos(ArithmosDomainNat)
}

// IsInt gets whether this type is that of integers.
func (t *RsType) IsInt() bool {
	return t.IsArithmos(ArithmosDomainInt)
}

// IsReal gets whether this type is that of real numbers.
func (t *RsType) IsReal() bool {
	return t.IsArithmos(ArithmosDomainReal)
}

// IsArithmos gets whether this type is an arithmetic one with the given domain.
func (t *RsType) IsArithmos(domain ArithmosDomain) bool {
	if t.kind != KindArithmos {
		return false
	}
	return t.arithmosDomain == domain
}

func (t *RsType) MarshalText() ([]byte, error) {
	if str, ok := t.tryString(); ok {
		return []byte(str), nil
	}

	return nil, TypeMarshalError{Type: *t}
}

func (t *RsType) String() string {
	if str, ok := t.tryString(); ok {
		return str
	}

	return fmt.Sprintf("%v", *t)
}

func (t *RsType) tryString() (string, bool) {
	// TODO: factor into primitive and non-primitive types
	if t.IsEmpty() {
		return "()", true
	}
	if t.IsNat() {
		return "nat", true
	}
	if t.IsInt() {
		return "int", true
	}
	if t.IsReal() {
		return "real", true
	}
	if t.IsEnum() {
		return "enum", true
	}

	return "", false
}

func (t *RsType) UnmarshalText(text []byte) error {
	t.kind = KindEmpty
	t.arithmosDomain = ArithmosDomainUnknown

	// TODO: factor into primitive and non-primitive types
	typeStr := string(bytes.ToLower(text))
	switch typeStr {
	case "()":
		t.kind = KindEmpty
		return nil
	case "enum":
		t.kind = KindEnum
		return nil
	case "nat":
		t.kind = KindArithmos
		t.arithmosDomain = ArithmosDomainNat
		return nil
	case "int":
		t.kind = KindArithmos
		t.arithmosDomain = ArithmosDomainInt
		return nil
	case "real":
		t.kind = KindArithmos
		t.arithmosDomain = ArithmosDomainReal
		return nil
	}

	return TypeUnmarshalError{TypeName: typeStr}
}

//
// Type constructors
//

// Empty captures an empty type.
func Empty() *RsType {
	return &RsType{
		kind:           KindEmpty,
		arithmosDomain: ArithmosDomainUnknown,
	}
}

// Enum captures an enumeration type.
func Enum() *RsType {
	return &RsType{
		kind:           KindEnum,
		arithmosDomain: ArithmosDomainUnknown,
	}
}

// Arithmos captures an arithmetic type.
func Arithmos(domain ArithmosDomain) *RsType {
	return &RsType{
		kind:           KindArithmos,
		arithmosDomain: domain,
	}
}

// Typeable represents types that have RoboStar types.
type Typeable interface {
	// Type gets the underlying RoboStar type of this item.
	Type() *RsType
}
