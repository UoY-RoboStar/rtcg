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

// Trace is the type of 'flat' forbidden-trace tests.
type Trace struct {
	Prefix []testlang.Event // Prefix is the sequence of events that must occur for the test to pass.
	Forbid testlang.Event   // Forbid is the event that must not occur after Prefix.
}

// New constructs a trace with forbidden event forbid and prefix after.
func New(forbid testlang.Event, after ...testlang.Event) Trace {
	return Trace{Prefix: after, Forbid: forbid}
}

func (t Trace) String() string {
	prefixStrs := make([]string, len(t.Prefix))

	for i, p := range t.Prefix {
		prefixStrs[i] = p.String()
	}

	prefixStr := strings.Join(prefixStrs, ", ")

	return fmt.Sprintf("<%s>!%s", prefixStr, &t.Forbid)
}
