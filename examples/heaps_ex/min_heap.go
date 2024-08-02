package heapsex

import (
	"fmt"

	"github.com/0x0FACED/go-collections/heaps"
)

func Example_MinHeap() {
	fmt.Println("Currenty: Example_MinHeap()")
	intCompare := func(a, b int) int {
		if a == b {
			return 0
		} else if a > b {
			return -1
		} else {
			return 1
		}
	}
	var minHeap heaps.Heap[int]
	minHeap = heaps.NewHeap(intCompare)

	for i := 1000; i >= 0; i-- {
		minHeap.Insert(i)
	}

	val, err := minHeap.Peek()
	if err != nil {
		return
	}
	fmt.Println("Root val is ", *val) // output: 0

	val, err = minHeap.Extract()
	if err != nil {
		return
	}
	fmt.Println("Root val Extract() is ", *val) // output: 0

	// current Peek must be 1
	val, _ = minHeap.Peek()
	fmt.Println("Peek() val is ", *val) // output: 1

	// lets delete all the elements
	for !minHeap.IsEmpty() {
		minHeap.Extract()
	}

	_, err = minHeap.Extract() // err == "empty"
	if err != nil {
		fmt.Println(err.Error()) // -> will print error
	}
}
