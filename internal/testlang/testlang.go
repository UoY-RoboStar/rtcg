// Package testlang implements the testing language.
package testlang

import (
	"encoding/json"
	"fmt"
	"io"
)

// Suite is a test suite, with each test having a name.
type Suite map[string]*Node

// ReadSuite reads a test suite from JSON in reader r.
func ReadSuite(r io.Reader) (Suite, error) {
	var suite Suite

	j := json.NewDecoder(r)
	if err := j.Decode(&suite); err != nil {
		return nil, fmt.Errorf("json decoding error for test suite: %w", err)
	}

	return suite, nil
}

// Write pretty-prints a test suite, as JSON, into writer w.
func (s *Suite) Write(w io.Writer) error {
	j := json.NewEncoder(w)
	j.SetIndent("", "\t")

	if err := j.Encode(s); err != nil {
		return fmt.Errorf("json encoding error for test suite: %w", err)
	}

	return nil
}
