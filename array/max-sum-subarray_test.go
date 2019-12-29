package array

import "testing"

func TestMaximumSumSubArray(t *testing.T) {
	var A = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	var expexted = []int{4, -1, 2, 1,}
	max := MaximumSumSubArray(A)

	if max != 6 {
		t.Errorf("wrong max sum, expected is sum of  %v got %d", expexted,  max)
	}
}
