package dynamicpropgramming

import "testing"

func TestLcsV1(t *testing.T) {
	A := []int{1, 2, 3, 2, 4, 1, 2}
	B := []int{2, 4, 3, 1, 2, 1, 2}

	expected := []int{2, 3, 2, 1, 2}

	lcs := LcsV1(A, B)

	if len(lcs) != len(expected) {
		t.Errorf("LcsV1(%v, %v) with result: %v", A, B, lcs)
		return
	}

	for i, v := range lcs {
		if v != expected[i] {
			t.Errorf("LcsV1(%v, %v) with result: %v", A, B, lcs)
			break
		}
	}
}
