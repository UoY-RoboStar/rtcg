package stm

import (
	"fmt"
	"io"

	"github.com/UoY-RoboStar/rtcg/internal/serial"
)

// Suite is a test suite, in state machine format.
type Suite struct {
	Types TypeMap         `json:"types,omitempty"` // Types is the unified type map for the channels.
	Tests map[string]*Stm `json:"tests,omitempty"` // Tests is the set of test STMs.
}

// ReadSuite reads a state machine suite from JSON in reader r.
func ReadSuite(r io.Reader) (*Suite, error) {
	var suite Suite

	if err := serial.ReadJSON(r, &suite); err != nil {
		return nil, fmt.Errorf("couldn't read state machine suite: %w", err)
	}

	return &suite, nil
}

// Write pretty-prints a state machine suite, as JSON, into writer w.
func (s *Suite) Write(w io.Writer) error {
	if err := serial.WriteJSON(w, s); err != nil {
		return fmt.Errorf("couldn't write state machine suite: %w", err)
	}

	return nil
}
