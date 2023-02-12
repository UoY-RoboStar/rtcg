package structure

// Set is a set, implemented as a map.
type Set[T comparable] map[T]bool

// Add adds each value in values to s.
func (s Set[T]) Add(values ...T) {
	for _, v := range values {
		s[v] = true
	}
}

// Values gets an unsorted slice of values present in s.
func (s Set[T]) Values() []T {
	values := make([]T, 0, len(s))

	for k, v := range s {
		if v {
			values = append(values, k)
		}
	}

	return values
}
