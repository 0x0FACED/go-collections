package trees

import "sync"

// bst - Binary Search Tree
type bst[T comparable] struct {
	root *node[T]

	mu sync.Mutex

	// comparator is the main func to compare two values of type T
	//
	// You have to write your compare func for your data type
	compare Comparator[T]
}

func NewBST[T comparable](compare Comparator[T]) *bst[T] {
	return &bst[T]{compare: compare}
}

func (bst *bst[T]) Insert(item T) {
	bst.mu.Lock()
	defer bst.mu.Unlock()

	bst.root = bst.insertHelper(bst.root, item)
}

func (bst *bst[T]) Delete(item T) error {
	bst.mu.Lock()
	defer bst.mu.Unlock()

	var err error
	bst.root, err = bst.deleteHelper(bst.root, item)
	return err
}

func (bst *bst[T]) Search(item T) (*T, error) {
	bst.mu.Lock()
	defer bst.mu.Unlock()

	return bst.searchHelper(bst.root, item)
}

func (bst *bst[T]) PreOrder() []T {
	bst.mu.Lock()
	defer bst.mu.Unlock()

	var items []T
	bst.preOrderHelper(bst.root, &items)
	return items
}

func (bst *bst[T]) InOrder() []T {
	bst.mu.Lock()
	defer bst.mu.Unlock()

	var items []T
	bst.inOrderHelper(bst.root, &items)
	return items
}

func (bst *bst[T]) PostOrder() []T {
	bst.mu.Lock()
	defer bst.mu.Unlock()

	var items []T
	bst.postOrderHelper(bst.root, &items)
	return items
}

func (bst *bst[T]) LevelOrder() []T {
	bst.mu.Lock()
	defer bst.mu.Unlock()

	return bst.levelOrderHelper()
}
