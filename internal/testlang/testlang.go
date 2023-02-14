// Package testlang implements the testing language.
package testlang

import (
	"fmt"
	"io"

	"github.com/UoY-RoboStar/rtcg/internal/serial"
)

// Suite is a test suite, with each test having a name.
type Suite map[string]*Node

// ReadSuite reads a test suite from JSON in reader r.
func ReadSuite(r io.Reader) (Suite, error) {
	var suite Suite

	if err := serial.ReadJSON(r, &suite); err != nil {
		return nil, fmt.Errorf("couldn't read test suite: %w", err)
	}

	return suite, nil
}

// Write pretty-prints a test suite, as JSON, into writer w.
func (s *Suite) Write(w io.Writer) error {
	if err := serial.WriteJSON(w, s); err != nil {
		return fmt.Errorf("couldn't write test suite: %w", err)
	}

	return nil
}
