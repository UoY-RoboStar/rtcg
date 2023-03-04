// Code generated by go-enum DO NOT EDIT.
// Version: 0.5.5
// Revision: b9e7d1ac24b2b7f6a5b451fa3d21706ffd8d79e2
// Build Date: 2023-01-30T01:49:43Z
// Built By: goreleaser

package cpp

import (
	"errors"
	"fmt"
	"strings"
)

const (
	// VariantAnimate is a Variant of type Animate.
	// Variant with a manually fed event loop.
	VariantAnimate Variant = iota
	// VariantRos is a Variant of type Ros.
	// Variant targeting ROS1 Noetic.
	VariantRos
)

var ErrInvalidVariant = errors.New("not a valid Variant")

const _VariantName = "animateros"

var _VariantMap = map[Variant]string{
	VariantAnimate: _VariantName[0:7],
	VariantRos:     _VariantName[7:10],
}

// String implements the Stringer interface.
func (x Variant) String() string {
	if str, ok := _VariantMap[x]; ok {
		return str
	}
	return fmt.Sprintf("Variant(%d)", x)
}

var _VariantValue = map[string]Variant{
	_VariantName[0:7]:                   VariantAnimate,
	strings.ToLower(_VariantName[0:7]):  VariantAnimate,
	_VariantName[7:10]:                  VariantRos,
	strings.ToLower(_VariantName[7:10]): VariantRos,
}

// ParseVariant attempts to convert a string to a Variant.
func ParseVariant(name string) (Variant, error) {
	if x, ok := _VariantValue[name]; ok {
		return x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := _VariantValue[strings.ToLower(name)]; ok {
		return x, nil
	}
	return Variant(0), fmt.Errorf("%s is %w", name, ErrInvalidVariant)
}

// MarshalText implements the text marshaller method.
func (x Variant) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *Variant) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseVariant(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
