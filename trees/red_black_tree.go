package trees

import (
	"fmt"
	"sync"

	gocollections "github.com/0x0FACED/go-collections"
)

// rbt - Red-Black Tree
type rbt[T comparable] struct {
	root *rbt_node[T]

	mu sync.Mutex

	compare Comparator[T]
}

func NewRBT[T comparable](compare Comparator[T]) *rbt[T] {
	return &rbt[T]{compare: compare}
}

func (rbt *rbt[T]) Insert(item T) {
	rbt.mu.Lock()
	defer rbt.mu.Unlock()

	newNode := rbt.insertHelper(rbt.root, item)

	rbt.fixInsert(newNode)
}

func (rbt *rbt[T]) Delete(item T) error {
	rbt.mu.Lock()
	defer rbt.mu.Unlock()

	node := rbt.searchHelper(rbt.root, item)
	if node == nil {
		return fmt.Errorf(gocollections.ErrNotFound)
	}
	rbt.deleteHelper(node)

	return nil
}

func (rbt *rbt[T]) Search(item T) (*T, error) {
	rbt.mu.Lock()
	defer rbt.mu.Unlock()

	node := rbt.searchHelper(rbt.root, item)
	if node != nil {
		return &node.val, nil
	}
	return nil, fmt.Errorf(gocollections.ErrNotFound)

}

func (rbt *rbt[T]) InOrder() []T {
	rbt.mu.Lock()
	defer rbt.mu.Unlock()

	var items []T
	rbt.inOrderHelper(rbt.root, &items)
	return items
}

func (rbt *rbt[T]) PreOrder() []T {
	rbt.mu.Lock()
	defer rbt.mu.Unlock()

	var items []T
	rbt.preOrderHelper(rbt.root, &items)
	return items
}

func (rbt *rbt[T]) PostOrder() []T {
	rbt.mu.Lock()
	defer rbt.mu.Unlock()

	var items []T
	rbt.postOrderHelper(rbt.root, &items)
	return items
}

func (rbt *rbt[T]) LevelOrder() []T {
	rbt.mu.Lock()
	defer rbt.mu.Unlock()

	return rbt.levelOrderHelper()
}
func (rbt *rbt[T]) PrintTree() {
	rbt.mu.Lock()
	defer rbt.mu.Unlock()

	rbt.printTree(rbt.root, "", true)
}
