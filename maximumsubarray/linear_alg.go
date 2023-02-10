package maximumsubarray

func LinearAlg(arr int[]) (int, int) {
	// pos of start and end of maximum subarray position
	s, e := 0, 0;
	// sum of maximum subarray
	sum := arr[0];

	// current found maximum subarray posistion
	i, j := 0, 1;
	// sum of current found maximum subarray
	curr_sum := arr[0];
	for j < len(arr) {
		if arr[i] < 0 && i != j {
			// When value of element i is negtive integer,
			// meaning the maximum subarray should exclue element i.
			i += 1
			curr_sum -= arr[i];
		} else {
			if sum < curr_sum {
				s, e = i, j;
				sum = curr_sum
			}
			// extends the range of current processed maximum subarray.
			j += 1
			continue;
		}
	}

	return s, e;
}
