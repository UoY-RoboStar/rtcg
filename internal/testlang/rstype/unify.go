package rstype

import "fmt"

// Unify tries to unify two types together.
func Unify(x, y *RsType) (*RsType, error) {
	if x == nil {
		return y, nil
	}
	if y == nil {
		return x, nil
	}
	if x.kind != y.kind {
		return nil, fmt.Errorf("%w: %s and %s", ErrUnifyDifferentKinds, x, y)
	}

	// Base now, we can just look at x's kind.
	if x.kind == KindArithmos {
		return Arithmos(MaxDomain(x.arithmosDomain, y.arithmosDomain)), nil
	}

	// The other kinds don't have other things we need to unify (yet):
	return x, nil
}
