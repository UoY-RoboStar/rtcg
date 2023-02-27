package structure

// TryOverMapValues maps mapper over all values of input, returning the resulting new map or error.
func TryOverMapValues[K comparable, T, U any](input map[K]T, mapper func(T) (U, error)) (map[K]U, error) {
	result := make(map[K]U, len(input))

	for key, v := range input {
		nv, err := mapper(v)
		if err != nil {
			return nil, err
		}

		result[key] = nv
	}

	return result, nil
}

// OverMap maps mapper over all key-value pairs of input, returning the resulting new map.
func OverMap[K, L comparable, T, U any](input map[K]T, mapper func(K, T) (L, U)) map[L]U {
	result := make(map[L]U, len(input))

	for k, v := range input {
		nk, nv := mapper(k, v)
		result[nk] = nv
	}

	return result
}
