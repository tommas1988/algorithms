package sort

import (
	"testing"
)

func TestHeapSort(t *testing.T) {
	test := []int{
		4, 1, 3, 2, 16, 9, 10, 14, 8, 7,
	}
	ascExpected := []int{
		1, 2, 3, 4, 7, 8, 9, 10, 14, 16,
	}
	descExpected := []int{
		16, 14, 10, 9, 8, 7, 4, 3, 2, 1,
	}
	unsorted := make([]int, len(test))
	copy(unsorted, test)

	HeapSort[int](test, func(e1, e2 int) int {
		return e1 - e2
	}, ASC)

	for i, e := range test {
		if e != ascExpected[i] {
			t.Errorf("HeapSort[int](%v, compare_func, ASC) with result: %v", unsorted, test)
			return
		}
	}

	// revert test array
	copy(test, unsorted)

	HeapSort[int](test, func(e1, e2 int) int {
		return e1 - e2
	}, DESC)

	for i, e := range test {
		if e != descExpected[i] {
			t.Errorf("HeapSort[int](%v, compare_func, DESC) with result: %v", unsorted, test)
			return
		}
	}
}
