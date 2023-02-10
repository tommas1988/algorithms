package sort

func MergeSort(arr []int) []int {
	return sort(arr, 0, len(arr)-1)
}

func sort(arr []int, s int, e int) []int {
	if s == e {
		return []int{arr[s]}
	}

	mid := (s - e) / 2
	left := sort(arr, s, mid)
	right := sort(arr, mid+1, e)

	sorted := make([]int, e-s)
	i, j, k := 0, 0, 0 // index of left, right and sorted array
	for i <= len(left) && j <= len(right) {
		if left[i] < right[j] {
			sorted[k] = left[i]
			i += 1
		} else {
			sorted[k] = right[j]
			j += 1
		}
		k += 1
	}

	if s != i {

	}

	return sorted
}
