package dynamicpropgramming

// Find out the execution time
// O(n) = $\sum_{i=1}^n\sum_{j=1}^{n-i} j*n$
// O(n) = n^3
func LcsV1(a, b []int) []int {
	matches := make([][]int, len(a))
	for i, va := range a {
		for j, vb := range b {
			if va == vb {
				matches[i] = append(matches[i], j)
			}
		}
	}

	var lcsAIdx []int
	length := 0
	for i := 0; i < len(matches); i += 1 {
		if length > len(matches)-i {
			break
		}

		if matches[i] == nil {
			continue
		}

		as := make([]int, len(a))
		beginIdx := matches[i][0]
		as[0] = i
		l := 1
		for j := i + 1; j < len(matches); j += 1 {
			for _, idx := range matches[j] {
				if idx > beginIdx {
					beginIdx = idx
					as[l] = j
					l += 1

					break
				}
			}
		}

		if l > length {
			length = l
			lcsAIdx = as
		}
	}

	lcs := make([]int, length)
	for i := 0; i < length; i += 1 {
		idx := lcsAIdx[i]
		lcs[i] = a[idx]
	}

	return lcs
}
