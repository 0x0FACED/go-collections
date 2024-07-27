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

// Super Uber Mega HARD to implement
//
// IMPORTANT: node that we want to delete = NODE
//
// So, we ha ve a bag of cases with delete operation:
//
// 1. If NODE have NO children -> just delete NODE AND if NODE.clr = black -> fixDelete with nil child and node.parent
//
// 2. If NODE have 1 child
func (rbt *rbt[T]) deleteHelper(node *rbt_node[T]) {
	var child, par *rbt_node[T]
	var color COLOR

	// if NO children
	if node.left == nil && node.right == nil {
		// if node.parent == nil -> we want to delete node with no children -> tree will be nil
		if node.parent == nil {
			rbt.root = nil
			return
		}

		// if out NODE is parent.left child
		if node.parent.left == node {
			node.parent.left = nil
		} else { // otherwise, parent.right
			node.parent.right = nil
		}

		// Important moment: if NODE is black -> his parent black too
		// So we have to call fixDelete with nil child and node.parent
		if node.clr == black {
			rbt.fixDelete(nil, node.parent)
		}

		// fast return
		return
	}

	// If NODE has 1 child: left or right
	if (node.left != nil && node.right == nil) || (node.left == nil && node.right != nil) {
		// link parent
		par = node.parent

		// we determine where child is located
		if node.left == nil {
			child = node.right
		} else {
			child = node.left
		}
		// save color
		color = node.clr

		// change child.parent from NODE to NODE.parent (par)
		child.parent = par

		// if parent == root
		if par == nil {
			rbt.root = child
		} else {
			if node == par.left {
				par.left = child
			} else {
				par.right = child
			}
		}
		// if NODE was black -> fixDelete with our child and parent
		if color == black {
			rbt.fixDelete(child, par)
		}
		return
	}

	// The last case: NODE has 2 children
	// We ALWAYS have to find highest element from LEFT subtree!
	// For example:
	// we have tree:
	//
	//								100 (root)
	//							  /            \
	//                         [95]             105
	//                        /    \          /     \
	//                      90      97      103      110
	//                    /    \           /
	//                  85      92       102
	//                 /  \    /  \
	//               80   87  91  (93) <- this is will be our new val of node with val 95
	//
	// And we want to delete node with val = 95
	// So, this node has 2 children
	// We have to go to left subtree and than go to the right
	// And we will find max element
	// In this example we will find val = 93
	// successor = 93
	// we recursively call deleteHelper with new node.val
	// We do it because we need to fix our tree if needed (color black)
	//
	// successor always have 0 or 1 child
	successor := rbt.findMinRight(node.left)
	node.val = successor.val
	rbt.deleteHelper(successor)
}

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
