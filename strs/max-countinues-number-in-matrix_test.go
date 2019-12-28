package strs

import "testing"

func TestFindMathPath(t *testing.T) {
	matrix := [][]int{
		{9, 9, 7},
		{9, 7, 2},
		{6, 9, 5},
		{9, 1, 2},
	}

	result := FindMaxPath(matrix)
	if result != "997952" {
		t.Errorf("wrong number, got %s", result)
	}
}
