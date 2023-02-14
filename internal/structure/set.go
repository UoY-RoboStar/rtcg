package structure

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Set is a set, implemented as a map.
type Set[T comparable] struct {
	contents map[T]bool
}

// NewSet constructs a new Set with the given values.
func NewSet[T comparable](values ...T) Set[T] {
	result := Set[T]{contents: map[T]bool{}}
	result.Add(values...)

	return result
}

// Add adds each value in values to s.
func (s *Set[T]) Add(values ...T) {
	for _, v := range values {
		s.contents[v] = true
	}
}

// Values gets an unsorted slice of values present in s.
func (s *Set[T]) Values() []T {
	values := make([]T, 0, len(s.contents))

	for k, v := range s.contents {
		if v {
			values = append(values, k)
		}
	}

	return values
}

// MarshalJSON marshals a Set as its Values slice.
func (s *Set[T]) MarshalJSON() ([]byte, error) {
	bs, err := json.Marshal(s.Values())
	if err != nil {
		return nil, fmt.Errorf("couldn't marshal values for set: %w", err)
	}

	return bs, nil
}

// UnmarshalJSON unmarshals a Set as if it were a Values slice, omitting duplicates.
func (s *Set[T]) UnmarshalJSON(bytes []byte) error {
	var values []T

	if err := json.Unmarshal(bytes, &values); err != nil {
		return fmt.Errorf("couldn't unmarshal values for set: %w", err)
	}

	s.contents = map[T]bool{}

	s.Add(values...)

	return nil
}

func (s *Set[T]) String() string {
	values := s.Values()
	valueStrs := make([]string, len(values))

	for i, v := range values {
		valueStrs[i] = fmt.Sprint(v)
	}

	return fmt.Sprintf("[%s]", strings.Join(valueStrs, ", "))
}
