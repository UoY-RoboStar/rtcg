package rstype

//go:generate go-enum --marshal -nocase

// ArithmosDomain is an enumeration of kinds of arithmetic type.
/*
ENUM(
unknown, // This is not arithmetic, or is unspecified.
nat,     // A natural number.
int,     // An integral number.
real,    // A real number.
)
*/
// Any type that is not of kind KindArithmos will return ArithmosKindUnknown here.
type ArithmosDomain uint8

// MaxDomain gets the maximum (least-specific) domain between lhs and rhs.
func MaxDomain(lhs, rhs ArithmosDomain) ArithmosDomain {
	// All the domain enums are in ascending order, so we can just do this:
	if lhs <= rhs {
		return rhs
	}

	return lhs
}
