package trees

import (
	"fmt"

	gocollections "github.com/0x0FACED/go-collections"
)

// If a == b -> return 0
//
// If a > b -> return 1
//
// If a < b -> return -1
//
// val `a` is the new item, val `b` is the tree's item
type Comparator[T comparable] func(a, b T) int

// BST - Binary Search Tree
type BST[T comparable] struct {
	root *node[T]

	// comparator is the main func to compare two values of type T
	//
	// You have to write your compare func for your data type
	compare Comparator[T]
}

func NewBST[T comparable](compare Comparator[T]) *BST[T] {
	return &BST[T]{compare: compare}
}

func (bst *BST[T]) Insert(item T) {
	bst.root = bst.insertHelper(bst.root, item)
}

func (bst *BST[T]) Delete(item T) error {
	panic("not implemented") // TODO: Implement
}

func (bst *BST[T]) Search(item T) (*T, error) {
	return bst.searchHelper(bst.root, item)
}

func (bst *BST[T]) InOrder() []T {
	panic("not implemented") // TODO: Implement
}

func (bst *BST[T]) insertHelper(curr *node[T], item T) *node[T] {
	if curr == nil {
		return &node[T]{val: item}
	}
	// if new item < curr.val (tree's val) than we go to left subtree
	// else -> our new value > curr.val -> go to right
	if bst.compare(item, curr.val) == -1 {
		curr.left = bst.insertHelper(curr.left, item)
	} else {
		curr.right = bst.insertHelper(curr.right, item)
	}
	return curr
}

func (bst *BST[T]) searchHelper(curr *node[T], item T) (*T, error) {
	if curr == nil {
		return nil, fmt.Errorf(gocollections.ErrNotFound)
	}

	// if a == b
	if bst.compare(item, curr.val) == 0 {
		return &curr.val, nil
	}

	// if a < b
	// a - is the new element (item), b -> element of tree (curr.val)
	if bst.compare(item, curr.val) == -1 {
		return bst.searchHelper(curr.left, item)
	}
	return bst.searchHelper(curr.right, item)
}
