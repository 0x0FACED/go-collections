package trees

import "sync"

type avl[T comparable] struct {
	root *avl_node[T]

	mu sync.RWMutex

	compare Comparator[T]
}

func (avl *avl[T]) Insert(item T) {
	avl.root = avl.insertHelper(avl.root, item)
}

func (avl *avl[T]) Delete(item T) error {
	return nil
}

func (avl *avl[T]) Search(item T) (*T, error) {
	val, err := avl.searchHelper(item)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (avl *avl[T]) InOrder() []T {

}
func (avl *avl[T]) PreOrder() []T {

}
func (avl *avl[T]) PostOrder() []T {

}

// LevelOrder uses Queue and bfs to traverse
func (avl *avl[T]) LevelOrder() []T {

}
