package trees

// rbt - Red-Black Tree
type rbt[T comparable] struct {
	root *rbt_node[T]

	compare Comparator[T]
}

func NewRBT[T comparable]() *rbt_node[T] {
	return &rbt_node[T]{}
}

func (rbt *rbt[T]) Insert(item T) {
	panic("not implemented") // TODO: Implement
}

func (rbt *rbt[T]) Delete(item T) error {
	panic("not implemented") // TODO: Implement
}

func (rbt *rbt[T]) Search(item T) (*T, error) {
	panic("not implemented") // TODO: Implement
}
