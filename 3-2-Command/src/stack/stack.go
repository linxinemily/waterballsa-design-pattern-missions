// Package stack https://github.com/surzia/go-store/blob/master/collections/stack/stack.go
package stack

type Stack[T interface{}] struct {
	array []T
}

func NewStack[T interface{}]() *Stack[T] {
	stack := &Stack[T]{}
	stack.array = []T{}

	return stack
}

// Push adds t to the top of the stack
func (s *Stack[T]) Push(t T) {
	s.array = append(s.array, t)
}

// Pop removes the top element from the stack
func (s *Stack[T]) Pop() (t T) {
	if s.IsEmpty() {
		return
	}

	item := s.array[len(s.array)-1]
	s.array = s.array[0 : len(s.array)-1]
	return item
}

// Size returns the size of the stack
func (s *Stack[T]) Size() int {
	return len(s.array)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.array) == 0
}
