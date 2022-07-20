package heap

type Comparator interface {
	compare(e1 interface{}, e2 interface{}) int
}

type Heap struct {
	elts       []interface{}
	comparator Comparator
}

func (h *Heap) SetElements(elts []interface{}) {
	h.elts = elts
}

func (h *Heap) MaxHeapify(idx int) {
	for {
		left := h.LeftChild(idx)
		right := h.RightChild(idx)

		largest := idx
		if h.comparator.compare(h.elts[largest], h.elts[left]) < 0 {
			largest = left
		}
		if h.comparator.compare(h.elts[largest], h.elts[right]) < 0 {
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

func (h *Heap) MinHeapify(idx int) {
	for {
		left := h.LeftChild(idx)
		right := h.RightChild(idx)

		smallest := idx
		if h.comparator.compare(h.elts[smallest], h.elts[left]) > 0 {
			smallest = left
		}
		if h.comparator.compare(h.elts[smallest], h.elts[right]) > 0 {
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

func (h *Heap) Parent(idx int) int {
	parentIdx := (idx + 1) >> 1
	return parentIdx - 1
}

func (h *Heap) LeftChild(idx int) int {
	childIdx := (idx + 1) << 1
	return childIdx - 1
}

func (h *Heap) RightChild(idx int) int {
	childIdx := ((idx + 1) << 1) + 1
	return childIdx - 1
}
