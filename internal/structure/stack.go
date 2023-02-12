package structure

// Stack is a(n inefficient) implementation of a stack.
type Stack[T any] struct {
	items []T
	count int
}

// Push pushes item onto the stack.
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
	s.count++
}

// Pop pops an item from s.
//
// It does not check to see if the stack is empty.
func (s *Stack[T]) Pop() T {
	item := s.items[s.count-1]
	s.items = s.items[:s.count-1]
	s.count--

	return item
}

// IsEmpty is true iff the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return s.count == 0
}

// Clear empties stack s.
func (s *Stack[T]) Clear() {
	s.items = []T{}
	s.count = 0
}
