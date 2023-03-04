// Code generated by go-enum DO NOT EDIT.
// Version: 0.5.5
// Revision: b9e7d1ac24b2b7f6a5b451fa3d21706ffd8d79e2
// Build Date: 2023-01-30T01:49:43Z
// Built By: goreleaser

package rstype

import (
	"errors"
	"fmt"
	"strings"
)

const (
	// ArithmosDomainUnknown is a ArithmosDomain of type Unknown.
	// This is not arithmetic, or is unspecified.
	ArithmosDomainUnknown ArithmosDomain = iota
	// ArithmosDomainNat is a ArithmosDomain of type Nat.
	// A natural number.
	ArithmosDomainNat
	// ArithmosDomainInt is a ArithmosDomain of type Int.
	// An integral number.
	ArithmosDomainInt
	// ArithmosDomainReal is a ArithmosDomain of type Real.
	// A real number.
	ArithmosDomainReal
)

var ErrInvalidArithmosDomain = errors.New("not a valid ArithmosDomain")

const _ArithmosDomainName = "unknownnatintreal"

var _ArithmosDomainMap = map[ArithmosDomain]string{
	ArithmosDomainUnknown: _ArithmosDomainName[0:7],
	ArithmosDomainNat:     _ArithmosDomainName[7:10],
	ArithmosDomainInt:     _ArithmosDomainName[10:13],
	ArithmosDomainReal:    _ArithmosDomainName[13:17],
}

// String implements the Stringer interface.
func (x ArithmosDomain) String() string {
	if str, ok := _ArithmosDomainMap[x]; ok {
		return str
	}
	return fmt.Sprintf("ArithmosDomain(%d)", x)
}

var _ArithmosDomainValue = map[string]ArithmosDomain{
	_ArithmosDomainName[0:7]:                    ArithmosDomainUnknown,
	strings.ToLower(_ArithmosDomainName[0:7]):   ArithmosDomainUnknown,
	_ArithmosDomainName[7:10]:                   ArithmosDomainNat,
	strings.ToLower(_ArithmosDomainName[7:10]):  ArithmosDomainNat,
	_ArithmosDomainName[10:13]:                  ArithmosDomainInt,
	strings.ToLower(_ArithmosDomainName[10:13]): ArithmosDomainInt,
	_ArithmosDomainName[13:17]:                  ArithmosDomainReal,
	strings.ToLower(_ArithmosDomainName[13:17]): ArithmosDomainReal,
}

// ParseArithmosDomain attempts to convert a string to a ArithmosDomain.
func ParseArithmosDomain(name string) (ArithmosDomain, error) {
	if x, ok := _ArithmosDomainValue[name]; ok {
		return x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := _ArithmosDomainValue[strings.ToLower(name)]; ok {
		return x, nil
	}
	return ArithmosDomain(0), fmt.Errorf("%s is %w", name, ErrInvalidArithmosDomain)
}

// MarshalText implements the text marshaller method.
func (x ArithmosDomain) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *ArithmosDomain) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseArithmosDomain(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
