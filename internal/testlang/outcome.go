package testlang

//go:generate go-enum --marshal -nocase

// Outcome is an enumeration of statuses during a test.
/*
ENUM(
inc, // Inconclusive status.
fail, // Failing status.
pass // Passing status.
)
*/
// These are the same statuses understood in the CSP testing theory.
type Outcome uint8

// Most of the status constants are autogenerated by go-enum; see status_enum.go

const (
	FirstOutcome = OutcomeInc  // FirstOutcome is the first status.
	LastStatus   = OutcomePass // LastStatus is the last status.

	NumStatus = uint8(LastStatus) + 1 // NumStatus is the number of valid status entries.
)