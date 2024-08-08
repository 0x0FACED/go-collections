package trees

import (
	"fmt"

	gocollections "github.com/0x0FACED/go-collections"
)

func (avl *avl[T]) insertHelper(curr *avl_node[T], item T) *avl_node[T] {
	if curr == nil {
		return &avl_node[T]{val: item}
	}

	if avl.compare(item, curr.val) < 0 {
		curr.left = avl.insertHelper(curr.left, item)
	} else {
		curr.right = avl.insertHelper(curr.right, item)
	}

	return curr
}

// TODO: Impl
func (avl *avl[T]) deleteHelper(curr *avl_node[T], item T) *avl_node[T] {
	if curr == nil {
		return &avl_node[T]{val: item}
	}

	if avl.compare(item, curr.val) < 0 {
		curr.left = avl.insertHelper(curr.left, item)
	} else if avl.compare(item, curr.val) > 0 {
		curr.right = avl.insertHelper(curr.right, item)
	} else {
		return nil
	}

	return curr
}

func (avl *avl[T]) searchHelper(curr *avl_node[T], item T) (*T, error) {
	if curr == nil {
		return nil, fmt.Errorf(gocollections.ErrNotFound)
	}

	// if a == b
	if avl.compare(item, curr.val) == 0 {
		return &curr.val, nil
	}

	// if a < b
	// a - is the new element (item), b -> element of tree (curr.val)
	if avl.compare(item, curr.val) == -1 {
		return avl.searchHelper(curr.left, item)
	}
	return avl.searchHelper(curr.right, item)
}
