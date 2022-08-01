package heap

type Heap[E any] struct {
	size        int
	elts        []E
	compareFunc func(E, E) int
}

func New[E any](elts []E, size int, compareFunc func(E, E) int) *Heap[E] {
	return &Heap[E]{
		size:        size,
		elts:        elts,
		compareFunc: compareFunc,
	}
}

func (h *Heap[E]) SetHeapArray(elts []E, size int) {
	h.size = size
	h.elts = elts
}

func (h *Heap[E]) MaxHeapify(idx int) {
	for {
		left := h.LeftChild(idx)
		right := h.RightChild(idx)

		largest := idx
		if left < h.size && h.compareFunc(h.elts[largest], h.elts[left]) < 0 {
			largest = left
		}
		if right < h.size && h.compareFunc(h.elts[largest], h.elts[right]) < 0 {
			largest = right
		}

		if largest == idx {
			return
		}

		// switch largest element and continue process with new idx
		h.elts[largest], h.elts[idx] = h.elts[idx], h.elts[largest]
		idx = largest
	}

}

func (h *Heap[E]) MinHeapify(idx int) {
	for {
		left := h.LeftChild(idx)
		right := h.RightChild(idx)

		smallest := idx
		if left < h.size && h.compareFunc(h.elts[smallest], h.elts[left]) > 0 {
			smallest = left
		}
		if right < h.size && h.compareFunc(h.elts[smallest], h.elts[right]) > 0 {
			smallest = right
		}

		if smallest == idx {
			return
		}

		// switch smallest element and continue process with new idx
		h.elts[smallest], h.elts[idx] = h.elts[idx], h.elts[smallest]
		idx = smallest
	}
}

func (h *Heap[E]) Parent(idx int) int {
	parentIdx := (idx + 1) >> 1
	return parentIdx - 1
}

func (h *Heap[E]) LeftChild(idx int) int {
	childIdx := (idx + 1) << 1
	return childIdx - 1
}

func (h *Heap[E]) RightChild(idx int) int {
	childIdx := (idx + 1) << 1
	return childIdx
}
