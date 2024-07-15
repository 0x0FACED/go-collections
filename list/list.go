package list

type node[T comparable] struct {
	val  T
	next *node[T]
}

type list[T comparable] interface {
	// Add adds val to the end of list
	Add(item T) error

	// Insert adds `item` to position `pos`. Similar to Set,
	// but there is a shift of the remaining part to the right, that is,
	// there is no replacement of the element
	Insert(item T, pos int) error

	// Remove removes `item` from tail of list
	RemoveLast() error

	// RemoveVal removes the first `item` of list equals to `item` and returns its `pos`
	RemoveVal(item T) (int, error)

	// RemoveAt removes `item` in position `pos`
	RemoveAt(pos int) error

	// Set sets new `item` for position `pos`
	Set(item T, pos int) error

	// Get returns `item` from position `pos`
	Get(pos int) (*T, error)

	// GetPosition returns `pos` of the first `item`
	GetPosition(item T) (int, error)

	// Size returns current size of list
	Size() int

	// Clear clears the list
	Clear() error

	// Contains check if the `item` exists in list and returns true. Returns false if not.
	Contains(item T) bool
}
