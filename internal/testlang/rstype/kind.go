package rstype

//go:generate go-enum --marshal -nocase

// Kind is an enumeration of primitive type kinds.
/*
ENUM(
empty,    // An entity that has no value, and therefore no type.
arithmos, // A number.
enum,     // An enumeration.
)
*/
// Every type has a corresponding Kind, but may recursively contain other types defined by
// other primitives.
type Kind uint8
