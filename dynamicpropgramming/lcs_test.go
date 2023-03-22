package dynamicpropgramming

import "testing"

var A = []int{1, 2, 3, 2, 4, 1, 2}
var B = []int{2, 4, 3, 1, 2, 1, 2}
var expected = []int{2, 3, 2, 1, 2}

func TestLcsV1(t *testing.T) {
	lcs := LcsV1(A, B)
	verifyResult("LcsV1", lcs, t)
}

func TestLcsV2(t *testing.T) {
	lcs := LcsV2(A, B)
	verifyResult("LcsV2", lcs, t)
}

func TestLcsV3(t *testing.T) {
	lcs := LcsV3(A, B)
	verifyResult("LcsV3", lcs, t)
}

func verifyResult(function string, lcs []int, t *testing.T) {
	if len(lcs) != len(expected) {
		t.Errorf("%s(%v, %v) with result: %v", function, A, B, lcs)
		return
	}

	for i, v := range lcs {
		if v != expected[i] {
			t.Errorf("%s(%v, %v) with result: %v", function, A, B, lcs)
			break
		}
	}
}
