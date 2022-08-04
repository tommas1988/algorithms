package queue

import (
	"github.com/tommas1988/algorithms/datastructure/heap"
)

const min_priority = -1

type Entry[T any] struct {
	priority int
	value    T
	index    int
}

type PriorityQueue[T any] struct {
	count   int
	entries []Entry[T]
	heap    *heap.Heap[Entry[T]]
}

func New[T any](size int) *PriorityQueue[T] {
	q := &PriorityQueue[T]{}
	q.entries = make([]Entry[T], size)
	q.heap = heap.New(q.entries, 0,
		func(e1, e2 Entry[T]) int {
			return e1.priority - e2.priority
		}, q.updateIndex)
	return q
}

func (q *PriorityQueue[T]) Enqueue(priority int, value T) *Entry[T] {
	entry := &q.entries[q.count]
	if q.count == 0 {
		entry.priority = priority
		entry.value = value
		entry.index = 0
	} else {
		entry.priority = min_priority
		entry.value = value
		entry.index = q.count

		q.increasePriority(q.count, priority)
	}

	q.count += 1
	q.heap.SetHeapArray(q.entries, q.count)
	return entry
}

func (q *PriorityQueue[T]) Dequeue() T {
	if q.count == 0 {
		panic("empty queue")
	}

	maxEntry := q.entries[0]
	q.count -= 1
	q.entries[0] = q.entries[q.count]
	q.heap.SetHeapArray(q.entries, q.count)
	q.heap.MaxHeapify(0)

	return maxEntry.value
}

func (q *PriorityQueue[T]) Peek() T {
	return q.entries[0].value
}

func (q *PriorityQueue[T]) Remove(idx int) {
	q.count -= 1
	q.heap.SetHeapArray(q.entries, q.count)
	last := q.count
	if idx != last {
		q.entries[idx] = q.entries[last]
		q.heap.MaxHeapify(idx)
	}
}

func (q *PriorityQueue[T]) updateIndex(newIdx int) {
	q.entries[newIdx].index = newIdx
}

func (q *PriorityQueue[T]) increasePriority(idx int, priority int) {
	for idx > 0 {
		p := q.heap.Parent(idx)
		if q.entries[p].priority < priority {
			q.entries[p], q.entries[idx] = q.entries[idx], q.entries[p]
			q.updateIndex(p)
			q.updateIndex(idx)
			idx = p
		} else {
			break
		}
	}
}
