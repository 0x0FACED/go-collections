package trees

import "fmt"

func (rbt *rbt[T]) insertHelper(curr *rbt_node[T], item T) *rbt_node[T] {
	newNode := &rbt_node[T]{val: item, clr: red}

	if rbt.root == nil {
		rbt.root = newNode
		rbt.root.clr = black
		return rbt.root
	}

	var parent *rbt_node[T]

	for curr != nil {
		parent = curr
		if rbt.compare(item, curr.val) < 0 {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}

	newNode.parent = parent
	if rbt.compare(item, parent.val) < 0 {
		parent.left = newNode
	} else {
		parent.right = newNode
	}
	return newNode
}

func (rbt *rbt[T]) fixInsert(curr *rbt_node[T]) {
	// while our newNode != root and his parent color == red
	for curr != rbt.root && curr.parent.clr == red {
		// if newNode's parent == child of his grandpa
		// in other words, if we are on the left
		//
		//			10 (newNode.parent.parent) 	(always black)
		//		   /
		//        9 (newNode.parent) 			(red right now)
		//       /
		//      8 (newNode) 					(red as his parent -> thats bad, we have to rotate the tree)
		//
		// newNode.parent.val = 9
		// newNode.parent.parent.val (grandpa) = 10
		// newNode.parent.parent.left.val = 9 (child of grandpa)
		// so we on the left subtree
		if curr.parent == curr.parent.parent.left {
			// y - we assign this variable a reference to curr.parent.parent.right
			// to avoid writing a lot of looooooooong text
			// BTW curr.parent.parent.right -> uncle of newNode's parent.
			// Look:
			//
			//			10
			//		   /  \
			//        9    12 (12 for example. this is curr.parent.parent.right, because curr = 8)
			//       /
			//      8
			y := curr.parent.parent.right
			// if
			if y != nil && y.clr == red {
				curr.parent.clr = black
				y.clr = black
				curr.parent.parent.clr = red
				curr = curr.parent.parent
			} else {
				if curr == curr.parent.right {
					curr = curr.parent
					rbt.rotateLeft(curr)
				}
				curr.parent.clr = black
				curr.parent.parent.clr = red
				rbt.rotateRight(curr.parent.parent)
			}
		} else {
			y := curr.parent.parent.left
			if y != nil && y.clr == red {
				curr.parent.clr = black
				y.clr = black
				curr.parent.parent.clr = red
				curr = curr.parent.parent
			} else {
				if curr == curr.parent.left {
					curr = curr.parent
					rbt.rotateRight(curr)
				}
				curr.parent.clr = black
				curr.parent.parent.clr = red
				rbt.rotateLeft(curr.parent.parent)
			}
		}
	}
	rbt.root.clr = black
}

func (rbt *rbt[T]) rotateLeft(x *rbt_node[T]) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		rbt.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func (rbt *rbt[T]) rotateRight(y *rbt_node[T]) {
	x := y.left
	y.left = x.right
	if x.right != nil {
		x.right.parent = y
	}
	x.parent = y.parent
	if y.parent == nil {
		rbt.root = x
	} else if y == y.parent.right {
		y.parent.right = x
	} else {
		y.parent.left = x
	}
	x.right = y
	y.parent = x
}

func (rbt *rbt[T]) deleteHelper(node *rbt_node[T]) {
	y := node
	y_original_clr := y.clr
	var x *rbt_node[T]
	if node.left == nil {
		x = node.right
		rbt.transplant(node, node.right)
	} else if node.right == nil {
		x = node.left
		rbt.transplant(node, node.left)
	} else {
		y = rbt.findMin(node.right)
		y_original_clr = y.clr
		x = y.right
		if y.parent == node {
			if x != nil {
				x.parent = y
			}
		} else {
			rbt.transplant(y, y.right)
			y.right = node.right
			y.right.parent = y
		}
		rbt.transplant(node, y)
		y.left = node.left
		y.left.parent = y
		y.clr = node.clr
	}
	if y_original_clr == black {
		rbt.fixDelete(x)
	}
}

// find the in-order successor
func (rbt *rbt[T]) findMin(node *rbt_node[T]) *rbt_node[T] {
	dummy := node
	for dummy.left != nil {
		dummy = dummy.left
	}
	return dummy
}

func (rbt *rbt[T]) transplant(u, v *rbt_node[T]) {
	if u.parent == nil {
		rbt.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}

func (rbt *rbt[T]) fixDelete(x *rbt_node[T]) {
	for x != rbt.root && (x != nil && x.clr == black) {
		if x == x.parent.left {
			w := x.parent.right
			if w.clr == red {
				w.clr = black
				x.parent.clr = red
				rbt.rotateLeft(x.parent)
				w = x.parent.right
			}
			if (w.left == nil || w.left.clr == black) && (w.right == nil || w.right.clr == black) {
				w.clr = red
				x = x.parent
			} else {
				if w.right == nil || w.right.clr == black {
					w.left.clr = black
					w.clr = red
					rbt.rotateRight(w)
					w = x.parent.right
				}
				w.clr = x.parent.clr
				x.parent.clr = black
				if w.right != nil {
					w.right.clr = black
				}
				rbt.rotateLeft(x.parent)
				x = rbt.root
			}
		} else {
			w := x.parent.left
			if w.clr == red {
				w.clr = black
				x.parent.clr = red
				rbt.rotateRight(x.parent)
				w = x.parent.left
			}
			if (w.left == nil || w.left.clr == black) && (w.right == nil || w.right.clr == black) {
				w.clr = red
				x = x.parent
			} else {
				if w.left == nil || w.left.clr == black {
					w.right.clr = black
					w.clr = red
					rbt.rotateLeft(w)
					w = x.parent.left
				}
				w.clr = x.parent.clr
				x.parent.clr = black
				if w.left != nil {
					w.left.clr = black
				}
				rbt.rotateRight(x.parent)
				x = rbt.root
			}
		}
	}
	if x != nil {
		x.clr = black
	}
}

func (rbt *rbt[T]) searchHelper(curr *rbt_node[T], item T) *rbt_node[T] {
	dummy := curr
	for dummy != nil {
		if rbt.compare(item, dummy.val) < 0 {
			dummy = dummy.left
		} else if rbt.compare(item, dummy.val) > 0 {
			dummy = dummy.right
		} else {
			return dummy
		}
	}
	return nil
}

func (rbt *rbt[T]) inOrderHelper(curr *rbt_node[T], items *[]T) {
	if curr == nil {
		return
	}

	rbt.inOrderHelper(curr.left, items)
	*items = append(*items, curr.val)
	rbt.inOrderHelper(curr.right, items)
}

func printTree[T comparable](node *rbt_node[T], indent string, last bool) {
	if node != nil {
		fmt.Print(indent)
		if last {
			fmt.Print("R----")
			indent += "     "
		} else {
			fmt.Print("L----")
			indent += "|    "
		}
		color := "RED"
		if node.clr == black {
			color = "BLACK"
		}
		fmt.Printf("%v (%s)\n", node.val, color)
		printTree(node.left, indent, false)
		printTree(node.right, indent, true)
	}
}
