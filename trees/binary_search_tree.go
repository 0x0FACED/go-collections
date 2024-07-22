package trees

import (
	"fmt"

	gocollections "github.com/0x0FACED/go-collections"
	"github.com/0x0FACED/go-collections/queue"
)

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

func (bst *bst[T]) preOrderHelper(curr *node[T], items *[]T) {
	if curr == nil {
		return
	}
	*items = append(*items, curr.val)
	bst.inOrderHelper(curr.left, items)
	bst.inOrderHelper(curr.right, items)
}

func (bst *bst[T]) inOrderHelper(curr *node[T], items *[]T) {
	if curr == nil {
		return
	}

	bst.inOrderHelper(curr.left, items)
	*items = append(*items, curr.val)
	bst.inOrderHelper(curr.right, items)
}

func (bst *bst[T]) postOrderHelper(curr *node[T], items *[]T) {
	if curr == nil {
		return
	}

	bst.inOrderHelper(curr.left, items)
	bst.inOrderHelper(curr.right, items)
	*items = append(*items, curr.val)
}

func (bst *bst[T]) levelOrderHelper() []T {
	q := queue.NewDynamicListQueue[node[T]]()
	items := make([]T, 0)
	q.Enqueue(*bst.root)
	for !q.IsEmpty() {
		child, err := q.Dequeue()
		if err != nil {
			return nil
		}
		items = append(items, child.val)
		if child.left != nil {
			q.Enqueue(*child.left)
		}
		if child.right != nil {
			q.Enqueue(*child.right)
		}
	}

	return items
}

func (bst *bst[T]) insertHelper(curr *node[T], item T) *node[T] {
	if curr == nil {
		return &node[T]{val: item}
	}
	// if new item < curr.val compare returns -1 (tree's val) than we go to left subtree
	// else -> our new value >= curr.val -> go to right
	if bst.compare(item, curr.val) == -1 {
		curr.left = bst.insertHelper(curr.left, item)
	} else {
		curr.right = bst.insertHelper(curr.right, item)
	}
	return curr
}

func (bst *bst[T]) deleteHelper(curr *node[T], item T) (*node[T], error) {
	if curr == nil {
		return curr, fmt.Errorf(gocollections.ErrNotFound)
	}

	res := bst.compare(item, curr.val)
	var err error
	// if res == 0 -> item found in tree
	if res == 0 {
		if curr.left == nil && curr.right == nil {
			return nil, nil
		} else if curr.left == nil {
			return curr.right, nil
		} else if curr.right == nil {
			return curr.left, nil
		}

		// lets find min element of right subtree
		rightMin := bst.findMin(curr.right)
		curr.val = rightMin.val
		curr.right, err = bst.deleteHelper(curr.right, rightMin.val)
		if err != nil {
			return curr, err
		}
	} else if res == -1 { // if item < curr.val -> go left
		curr.left, err = bst.deleteHelper(curr.left, item)
		// if err != nil -> item not found, return err (gocollections.NotFound)
		if err != nil {
			return curr, err
		}
	} else { // else item > curr.val -> go right
		curr.right, err = bst.deleteHelper(curr.right, item)
		// if err != nil -> item not found, return err (gocollections.NotFound)
		if err != nil {
			return curr, err
		}
	}

	return curr, nil
}

func (bst *bst[T]) findMin(n *node[T]) *node[T] {
	for n.left != nil {
		n = n.left
	}
	return n
}

func (bst *bst[T]) searchHelper(curr *node[T], item T) (*T, error) {
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
