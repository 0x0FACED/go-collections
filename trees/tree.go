package trees

// ==========================================================================================

type COLOR string

const (
	// red color for node
	red = "RED"

	// black color for node
	black = "BLACK"
)

// ==========================================================================================

// If a == b -> return 0
//
// If a > b -> return 1
//
// If a < b -> return -1
//
// val `a` is the new item, val `b` is the tree's item
type Comparator[T comparable] func(a, b T) int

// ==========================================================================================

// rbt_node is the node of Red-Black Tree.
//
// # val
// -> stores the val of type T
//
// # clr 		-> color of node ("RED" or "BLACK")
//
// # left 	-> ptr to left child - subtree
//
// # right	-> ptr to right child - subtree
//
// # parent	-> prt to parent of this node - prev node
type rbt_node[T comparable] struct {
	val T

	// COLOR is alias for string
	clr    COLOR
	left   *rbt_node[T]
	right  *rbt_node[T]
	parent *rbt_node[T]
}

// ==========================================================================================

// traversal is the interface that has 3 methods to impl:
//
// InOrder(), PreOrder(), PostOrder()
//
// BST (Binary Search Tree) has these implementations
type traversal[T comparable] interface {
	InOrder() []T
	PreOrder() []T
	PostOrder() []T

	// LevelOrder uses Queue and bfs to traverse
	LevelOrder() []T
}

// ==========================================================================================

// node of tree
type node[T comparable] struct {
	val T

	right *node[T]
	left  *node[T]
}

// ==========================================================================================

// TraversalTree is the interface that has common tree operations (Insert, Delete, Search)
//
// and have 3 methods of traversal: InOrder, PreOrder, PostOrder
type TraversalTree[T comparable] interface {
	traversal[T]
	Tree[T]
}

// ==========================================================================================

// Tree interface has common methods of tree:
//
//	compare := func(a, b int) int {
//		if a == b {
//			return 0
//		} else if a < b {
//			return -1
//		} else {
//			return 1
//		}
//	}
//
//	var tr trees.Tree[int]
//
//	// you don't have to specify the generic type here -> NewBST`[TYPE]`
//	// bcz you use your custom compare with specific type
//	tr = trees.NewBST(compare)
//	tr.Insert(20)
//	tr.Insert(50)
//	var err error
//	err = tr.Delete(50)
//	val, err = tr.Search(20)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	fmt.Printf("Val: %v", *val)
type Tree[T comparable] interface {
	// Insert most common method of tree
	//
	// That method just adds element to tree
	Insert(item T)

	// Delete deletes element `item` arg
	//
	// if there is no item in tree -> returns err
	//
	// else -> nil
	Delete(item T) error

	// Search is searching `item` in tree
	//
	// if there is no item -> val = nil, err != nil
	//
	// else -> val != nil, err = nil
	Search(item T) (*T, error)
}
