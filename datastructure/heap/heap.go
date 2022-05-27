package heap

type element struct {
	key   int
	value interface{}
}

type heap struct {
	keys []int
	size int
}
type MaxHeap heap
type MinHeap heap

func NewMaxHeap(size int) *MaxHeap {
	return &MaxHeap{
		make([]int, size),
		0,
	}
}

func NewMinHeap(size int) *MinHeap {
	return &MinHeap{
		make([]int, size),
		0,
	}
}

func (h *heap) Insert(key int) {

}

func (h *heap) Delete(key int) bool {

}
