package trees

// bst - Binary Search Tree
type bst[T comparable] struct {
	root *node[T]

	// comparator is the main func to compare two values of type T
	//
	// You have to write your compare func for your data type
	compare Comparator[T]
}

func NewBST[T comparable](compare Comparator[T]) *bst[T] {
	return &bst[T]{compare: compare}
}

func (bst *bst[T]) Insert(item T) {
	bst.root = bst.insertHelper(bst.root, item)
}

func (bst *bst[T]) Delete(item T) error {
	var err error
	bst.root, err = bst.deleteHelper(bst.root, item)
	return err
}

func (bst *bst[T]) Search(item T) (*T, error) {
	return bst.searchHelper(bst.root, item)
}

func (bst *bst[T]) PreOrder() []T {
	var items []T
	bst.preOrderHelper(bst.root, &items)
	return items
}

func (bst *bst[T]) InOrder() []T {
	var items []T
	bst.inOrderHelper(bst.root, &items)
	return items
}

func (bst *bst[T]) PostOrder() []T {
	var items []T
	bst.postOrderHelper(bst.root, &items)
	return items
}

func (bst *bst[T]) LevelOrder() []T {
	return bst.levelOrderHelper()
}
