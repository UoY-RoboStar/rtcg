// Package trace concerns the ingestion of tests as forbidden trace.
//
// The format of a forbidden trace input file is a CSV file where each line is a separate trace (test) and each
// cell is an event; the last event in
package trace

import (
	"fmt"
	"strings"

	"github.com/UoY-RoboStar/rtcg/internal/testlang"
)

// Trace is the type of traces.
type Trace []testlang.Event // Events is the list of events in this trace.

func (t Trace) String() string {
	eventStrs := make([]string, len(t))

	for i, e := range t {
		eventStrs[i] = e.String()
	}

	return fmt.Sprintf("<%s>", strings.Join(eventStrs, ", "))
}

// Forbidden is the type of 'flat' forbidden-trace tests.
type Forbidden struct {
	Prefix Trace          // Prefix is the sequence of events that must occur for the test to pass.
	Forbid testlang.Event // Forbid is the event that must not occur after Prefix.
}

// New constructs a trace with forbidden event forbid and prefix after.
func New(forbid testlang.Event, after ...testlang.Event) Forbidden {
	return Forbidden{Prefix: after, Forbid: forbid}
}

func (t Forbidden) String() string {
	return fmt.Sprintf("%s!%s", &t.Prefix, &t.Forbid)
}
