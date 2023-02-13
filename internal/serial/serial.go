// Package serial contains helper functions for serialising and deserialising rtcg intermediate structures.
package serial

import (
	"encoding/json"
	"fmt"
	"io"
)

// ReadJSON reads a JSON representation from reader into dest.
func ReadJSON(reader io.Reader, dest any) error {
	j := json.NewDecoder(reader)

	if err := j.Decode(dest); err != nil {
		return fmt.Errorf("JSON read error: %w", err)
	}

	return nil
}

// WriteJSON writes a JSON representation for value into writer.
func WriteJSON(writer io.Writer, value any) error {
	j := json.NewEncoder(writer)
	j.SetIndent("", "\t")

	if err := j.Encode(value); err != nil {
		return fmt.Errorf("JSON write error: %w", err)
	}

	return nil
}
