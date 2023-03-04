package rstype

import "fmt"

// Unify tries to unify two types together.
func Unify(lhs, rhs *RsType) (*RsType, error) {
	switch {
	case lhs == nil:
		return rhs, nil
	case rhs == nil:
		return lhs, nil
	case lhs.kind != rhs.kind:
		return nil, fmt.Errorf("%w: %s and %s", ErrUnifyDifferentKinds, lhs, rhs)
	case lhs.kind == KindArithmos:
		// Arithmos types need to be unified on their domain, which is currently totally ordered.
		return Arithmos(MaxDomain(lhs.arithmosDomain, rhs.arithmosDomain)), nil
	default:
		// The other kinds don't have other things we need to unify (yet):
		return lhs, nil
	}
}
