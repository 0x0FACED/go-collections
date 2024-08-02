package heaps

type Comparator[T any] func(a, b T) int

// Standard intfc for heap
// To make MAX heap -> use compare like this:
//
//	func compare(a, b int) int {
//		if a == b {
//			return 0
//		} else if a < b {
//			return -1
//		} else {
//			return 1
//		}
//	}
//
// To make MIN heap -> use compare like this:
//
//	func compare(a, b int) int {
//		if a == b {
//			return 0
//		} else if a > b {
//			return -1
//		} else {
//			return 1
//		}
//	}
type Heap[T comparable] interface {
	// Insert adds element to heap
	Insert(item T)

	// Extract returns MAX item in Max-Heap and deletes it
	// and MIN in Min-Heap
	Extract() (*T, error)

	// Peek returns MAX item in Max-Heap
	// and MIN in Min-Heap
	Peek() (*T, error)

	// Size returns size of heap
	Size() int

	// IsEmpty() returns true is heap is empty, otherwise -> false
	IsEmpty() bool
}

// helper functions to get parent, left and right children indices
func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return i*2 + 1
}

func right(i int) int {
	return i*2 + 2
}
