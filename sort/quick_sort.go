package sort

func QuickSort(arr []int) {
	doQuickSort(arr, 0, len(arr)-1)
}

func doQuickSort(arr []int, s int, e int) {
	if s >= e {
		return
	}

	i := partition(arr, s, e)
	doQuickSort(arr, s, i-1)
	doQuickSort(arr, i+1, e)
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
