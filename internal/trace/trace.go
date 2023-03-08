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

// New constructs a Trace with the given events.
func New(events ...testlang.Event) Trace {
	return events
}

func (t Trace) String() string {
	eventStrs := make([]string, len(t))

	for i, e := range t {
		eventStrs[i] = e.String()
	}

	return fmt.Sprintf("<%s>", strings.Join(eventStrs, ", "))
}

// Equals checks whether t is equal to other.
func (t Trace) Equals(other Trace) bool {
	return len(t) == len(other) && t.IsPrefixOf(other)
}

func (t Trace) IsPrefixOf(other Trace) bool {
	// Can't be a prefix if we're bigger than the other trace
	if len(other) < len(t) {
		return false
	}

	for i, e := range t {
		if !e.Equals(other[i]) {
			return false
		}
	}

	return true
}

// Forbid constructs a Forbidden trace test with passing trace t and forbidden event forbid.
func (t Trace) Forbid(forbid testlang.Event) Forbidden {
	return t.ForbidWithName(forbid, "")
}

// ForbidWithName constructs a Forbidden with passing trace t, forbidden event forbid, and name.
func (t Trace) ForbidWithName(forbid testlang.Event, name string) Forbidden {
	return Forbidden{Name: name, Prefix: t, Forbid: forbid}
}

// Forbidden is the type of 'flat' forbidden-trace tests.
type Forbidden struct {
	Name   string         // Name is an optional name for the test.
	Prefix Trace          // Prefix is the sequence of events that must occur for the test to pass.
	Forbid testlang.Event // Forbid is the event that must not occur after Prefix.
}

// String gets a string representation of a forbidden-trace test.
func (t Forbidden) String() string {
	var nameTag string

	if t.Name != "" {
		nameTag = t.Name + ":"
	}

	return fmt.Sprintf("%s%s!%s", nameTag, &t.Prefix, &t.Forbid)
}

// Name assigns a systematic name to each trace in traces.
func Name(traces []Forbidden) map[string]Forbidden {
	result := make(map[string]Forbidden, len(traces))

	for i, t := range traces {
		name := fmt.Sprintf("test%d", i)
		result[name] = t
	}

	return result
}
