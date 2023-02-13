package sort

import "math/rand"

/**
 * Best with O(nlgn), and worst with O(n^2)
 */

func QuickSort(arr []int) {
	doQuickSort(arr, 0, len(arr)-1, false)
}

func RandomizedQuickSort(arr []int) {
	doQuickSort(arr, 0, len(arr)-1, true)
}

func doQuickSort(arr []int, s int, e int, randomize bool) {
	if s >= e {
		return
	}

	i := partition(arr, s, e)
	doQuickSort(arr, s, i-1, randomize)
	doQuickSort(arr, i+1, e, randomize)
}

// this version with less compare statements in each loop
func partition(arr []int, s int, e int) int {
	pivot := arr[e]
	j := s - 1 // refer to the last element of current left part array
	for i := s; i < e; i += 1 {
		if arr[i] < pivot {
			j += 1
			if j != i {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}

	j += 1
	arr[j], arr[e] = arr[e], arr[j]

	return j
}

func partitionV1(arr []int, s int, e int) int {
	pivot := arr[e]
	j := -1 // refer to the first element of current right part array
	for i := s; i < e; i += 1 {
		if arr[i] >= pivot && j == -1 {
			j = i
		} else if arr[i] < pivot && j != -1 {
			arr[i], arr[j] = arr[j], arr[i]
			j += 1
		}
	}

	arr[j], arr[e] = arr[e], arr[j]

	return j
}

// Randomize pivot element
func randomizedPartition(arr []int, s int, e int) int {
	i := rand.Intn(e-s+1) + s
	// switch random and end element
	arr[i], arr[e] = arr[e], arr[i]
	return partition(arr, s, e)
}
