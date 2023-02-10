package sort

import "testing"

func TestInsertionSort(t *testing.T) {
	test := []int{
		4, 1, 3, 2, 16, 9, 10, 14, 8, 7,
	}
	expected := []int{
		1, 2, 3, 4, 7, 8, 9, 10, 14, 16,
	}
	unsorted := make([]int, len(test))
	copy(unsorted, test)

	InsertionSort(test)

	for i, e := range test {
		if e != expected[i] {
			t.Errorf("InsertionSort(%v) with result: %v", unsorted, test)
			return
		}
	}
}
