package list

// dnode is a double node - node with next and prev ptrs
type dnode[T comparable] struct {
	val T

	next *dnode[T]
	prev *dnode[T]
}

// node is a struct of list to store val and ptr to next node
type node[T comparable] struct {
	val  T
	next *node[T]
}

// Common List interface with common operations
type List[T comparable] interface {
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

	// Get returns the last item of list
	GetLast() (*T, error)

	// GetPosition returns `pos` of the first `item`
	GetPosition(item T) (int, error)

	// Size returns current size of list
	Size() int

	// Clear clears the list
	Clear() error

	// Contains check if the `item` exists in list and returns true. Returns false if not.
	Contains(item T) bool
}

// Common List with List operations but includes Sort operation
type ListSort[T comparable] interface {
	List[T]
	sorter[T]
}

// Interface for lists with Sort operation
// cant use without List
type sorter[T comparable] interface {
	// Sort sorts list with unique compare method written by user
	Sort(compare Comparator[T], sortType int) error
}
