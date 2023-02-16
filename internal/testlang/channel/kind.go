package channel

//go:generate go-enum --marshal -nocase

// Kind is an enumeration of communication kinds.
/*
ENUM(
in,   // An event input.
out,  // An event output.
call, // An operation call.
)
*/
// These correspond to the input and output directions of RoboChart events, as well as operation calls.
type Kind uint8
