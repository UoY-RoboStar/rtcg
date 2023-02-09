package trace

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"

	"github.com/UoY-RoboStar/rtcg/internal/testlang"
)

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

	// Keep this as nil if there is no prefix, for normalisation.
	if 0 < prefixLen {
		trace.Prefix = make([]testlang.Event, prefixLen)
	}

	for cellIndex, cell := range row {
		ptr := selectEvent(cellIndex, prefixLen, &trace)

		if err := ptr.UnmarshalText([]byte(cell)); err != nil {
			return trace, fmt.Errorf("couldn't parse trace cell: %w", err)
		}
	}

	return trace, nil
}

func selectEvent(index int, prefixLen int, trace *Trace) *testlang.Event {
	if index == prefixLen {
		return &trace.Forbid
	}

	return &trace.Prefix[index]
}

var ErrNeedOneEvent = errors.New("each trace must have at least one event")
