// Package trace concerns the ingestion of tests as forbidden trace.
//
// The format of a forbidden trace input file is a CSV file where each line is a separate trace (test) and each
// cell is an event; the last event in
package trace

import (
	"encoding/csv"
	"errors"
	"io"
	"rtcg/internal/testlang"
)

// Trace is the type of 'flat' forbidden-trace tests.
type Trace struct {
	Prefix []testlang.Event // Prefix is the sequence of events that must occur for the test to pass.
	Forbid testlang.Event   // Forbid is the event that must not occur after Prefix.
}

// Read reads from r a list of traces.
func Read(r io.Reader) ([]Trace, error) {
	cr := newReader(r)

	var traces []Trace
	for {
		row, err := cr.Read()
		if errors.Is(err, io.EOF) {
			return traces, nil
		}
		if err != nil {
			return nil, err
		}

		trace, err := parseRow(row)
		if err != nil {
			return nil, err
		}

		traces = append(traces, trace)
	}
}

func newReader(r io.Reader) *csv.Reader {
	cr := csv.NewReader(r)
	cr.FieldsPerRecord = -1
	cr.Comment = '#'
	cr.TrimLeadingSpace = true
	return cr
}

func parseRow(row []string) (Trace, error) {
	var trace Trace

	prefixLen := len(row) - 1
	if prefixLen < 0 {
		return trace, ErrNeedOneEvent
	}

	trace.Prefix = make([]testlang.Event, prefixLen)
	for i, ev := range row {
		var ptr *testlang.Event
		if i == prefixLen {
			ptr = &trace.Forbid
		} else {
			ptr = &trace.Prefix[i]
		}
		if err := ptr.UnmarshalText([]byte(ev)); err != nil {
			return trace, err
		}
	}
	return trace, nil
}

var ErrNeedOneEvent = errors.New("each trace must have at least one event")
