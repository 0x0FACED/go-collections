package stack

type Stack[T comparable] interface {
	// Push adds the element to stack
	Push(item T)

	// Pop removes the last added element and returns it
	Pop() (*T, error)

	// Peek returns the last added element, but doesnt delete
	Peek() (*T, error)

	// Size returns size of stack
	Size() int

	// IsEmpty returns true if stack i empty, otherwise false
	IsEmpty() bool
}
