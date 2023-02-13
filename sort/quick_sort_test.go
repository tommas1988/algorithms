package sort

import "testing"

func TestQuickSort(t *testing.T) {
	test := []int{
		4, 1, 3, 2, 16, 9, 10, 14, 8, 7,
	}
	expected := []int{
		1, 2, 3, 4, 7, 8, 9, 10, 14, 16,
	}
	unsorted := make([]int, len(test))
	copy(unsorted, test)

	QuickSort(test)

	for i, e := range test {
		if e != expected[i] {
			t.Errorf("QuickSort(%v) with result: %v", unsorted, test)
			return
		}
	}
}

func TestRamdomizedQuickSort(t *testing.T) {
	test := []int{
		4, 1, 3, 2, 16, 9, 10, 14, 8, 7,
	}
	expected := []int{
		1, 2, 3, 4, 7, 8, 9, 10, 14, 16,
	}
	unsorted := make([]int, len(test))
	copy(unsorted, test)

	RandomizedQuickSort(test)

	for i, e := range test {
		if e != expected[i] {
			t.Errorf("RandomizedQuickSort(%v) with result: %v", unsorted, test)
			return
		}
	}
}

func TestPartition(t *testing.T) {
	test := []int{
		4, 1, 3, 2, 16, 7, 10, 14, 8, 9,
	}
	expectedResult := 6
	expectedArray := []int{
		4, 1, 3, 2, 7, 8, 9, 14, 16, 10,
	}
	unsorted := make([]int, len(test))
	copy(unsorted, test)

	result := partition(test, 0, len(test)-1)
	if result != expectedResult {
		t.Errorf("partition(%v) = %d", test, result)
	}

	for i, e := range test {
		if e != expectedArray[i] {
			t.Errorf("partition(%v) with result array: %v", unsorted, test)
			return
		}
	}
}

func TestPartitionV1(t *testing.T) {
	test := []int{
		4, 1, 3, 2, 16, 7, 10, 14, 8, 9,
	}
	expectedResult := 6
	expectedArray := []int{
		4, 1, 3, 2, 7, 8, 9, 14, 16, 10,
	}
	unsorted := make([]int, len(test))
	copy(unsorted, test)

	result := partitionV1(test, 0, len(test)-1)
	if result != expectedResult {
		t.Errorf("partitionV1(%v) = %d", test, result)
	}

	for i, e := range test {
		if e != expectedArray[i] {
			t.Errorf("partitionV1(%v) with result array: %v", unsorted, test)
			return
		}
	}
}
