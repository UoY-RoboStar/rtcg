// Package trace concerns the ingestion of tests as forbidden trace.
//
// The format of a forbidden trace input file is a CSV file where each line is a separate trace (test) and each
// cell is an event; the last event in
package trace

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"

	"github.com/UoY-RoboStar/rtcg/internal/testlang"
)

// Trace is the type of 'flat' forbidden-trace tests.
type Trace struct {
	Prefix []testlang.Event // Prefix is the sequence of events that must occur for the test to pass.
	Forbid testlang.Event   // Forbid is the event that must not occur after Prefix.
}

// Expand expands a single Trace into a test, tagged throughout with name.
func (t *Trace) Expand(name string) *testlang.Node {
	// Work backwards through the trace, building the tree from the failure.
	n := testlang.Pass(t.Forbid, testlang.Fail().From(name)).From(name)
	for i := len(t.Prefix) - 1; 0 <= i; i-- {
		n = testlang.Inc(t.Prefix[i], n).From(name)
	}

	return &n
}

// ExpandAll expands a list of traces to a systematically-named, non-factorised test suite.
func ExpandAll(traces []Trace) testlang.Suite {
	suite := make(testlang.Suite)

	for i, tr := range traces {
		name := fmt.Sprintf("test%d", i)
		suite[name] = tr.Expand(name)
	}

	return suite
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
			return nil, fmt.Errorf("error reading trace row: %w", err)
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

	for cellIndex, cell := range row {
		ptr := selectEvent(cellIndex, prefixLen, trace)

		if err := ptr.UnmarshalText([]byte(cell)); err != nil {
			return trace, fmt.Errorf("couldn't parse trace cell: %w", err)
		}
	}

	return trace, nil
}

func selectEvent(index int, prefixLen int, trace Trace) *testlang.Event {
	if index == prefixLen {
		return &trace.Forbid
	}

	return &trace.Prefix[index]
}

var ErrNeedOneEvent = errors.New("each trace must have at least one event")
