package trees

// rbt - Red-Black Tree
type rbt[T comparable] struct {
	root *rbt_node[T]

	compare Comparator[T]
}

func NewRBT[T comparable](compare Comparator[T]) *rbt[T] {
	return &rbt[T]{compare: compare}
}

func (rbt *rbt[T]) Insert(item T) {
	newNode := rbt.insertHelper(rbt.root, item, nil)

	rbt.fixInsert(newNode)
}

func (rbt *rbt[T]) Delete(item T) error {
	panic("not implemented") // TODO: Implement
}

func (rbt *rbt[T]) Search(item T) (*T, error) {
	panic("not implemented") // TODO: Implement
}

func (rbt *rbt[T]) InOrder() []T {
	var items []T
	rbt.inOrderHelper(rbt.root, &items)
	return items
}

func (rbt *rbt[T]) PrintTree() {
	printTree(rbt.root, "", true)
}
