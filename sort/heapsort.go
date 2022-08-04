package sort

import (
	"github.com/tommas1988/algorithms/datastructure/heap"
)

type Order int

const (
	ASC Order = iota
	DESC
)

func HeapSort[E any](elts []E, compareFunc func(E, E) int, order Order) {
	size := len(elts)
	h := heap.New[E](elts, size, compareFunc, nil)

	var heapify func(int)
	if order == DESC {
		heapify = h.MinHeapify
	} else {
		heapify = h.MaxHeapify
	}

	// build heap
	// heapify element that with children from bottom to top
	for i := h.Parent(size - 1); i >= 0; i-- {
		heapify(i)
	}

	for size > 1 {
		// switch last and largest elements and heapify with first element
		elts[size-1], elts[0] = elts[0], elts[size-1]
		size -= 1
		h.SetHeapArray(elts, size)
		heapify(0)
	}
}
