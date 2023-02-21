package testlang

//go:generate go-enum --marshal -nocase

// Outcome is an enumeration of outcomes during a test.
/*
ENUM(
unset, // No outcome recorded yet.
inc,   // Inconclusive outcome.
fail,  // Failing outcome.
pass   // Passing outcome.
)
*/
// These are the same statuses understood in the CSP testing theory.
type Outcome uint8
