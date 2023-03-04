package rstype

import (
	"errors"
	"fmt"
)

// TypeMarshalError is an error returned when we try to unmarshal an unknown type name.
type TypeMarshalError struct {
	Type RsType // RsType is the (unknown) type.
}

func (e TypeMarshalError) Error() string {
	return fmt.Sprintf("invalid type: %v", e.Type)
}

// TypeUnmarshalError is an error returned when we try to unmarshal an unknown type name.
type TypeUnmarshalError struct {
	TypeName string // TypeName is the (unknown) type name.
}

func (e TypeUnmarshalError) Error() string {
	return fmt.Sprintf("unknown type name: %s", e.TypeName)
}

// ErrUnifyDifferentKinds occurs when we try to unify two types with different kinds.
var ErrUnifyDifferentKinds = errors.New("can't unify types of different kinds")
