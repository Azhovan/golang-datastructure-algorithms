package array

import "math"

// Given an array and an input integer X,
//write a program to find the subarray of length X with least/minimum average

// Input:  arr[] = {3, 7, 90, 20, 10, 50, 40}, k = 3
//Output: Subarray between indexes 3 and 5
//The subarray {20, 10, 50} has the least average
//among all subarrays of size 3.
//
//Input:  arr[] = {3, 7, 5, 20, -10, 0, 12}, k = 2
//Output: Subarray between [4, 5] has minimum average

// Time complexity: O(n)
// efficient way
func MinAverageSubarray2(A []int, k int) []int {
	if len(A) < k {
		return []int{}
	}
	minAvg := math.MaxInt64
	index:= 0

	// find the first window
	windowAvg := 0
	for i := 0; i < k; i++ {
		windowAvg += A[i]
	}

	// move window
	for i := 1; i < len(A) && i+k <= len(A); i++ {
		windowAvg += A[i+k-1] - A[i-1]
		if windowAvg < minAvg {
			minAvg = windowAvg
			index = i
		}
	}
	return A[index:index+k]
}

// Time complexity : O(nk)
// this algorithm is not bad at all
func MinAverageSubarray1(A []int, k int) []int {
	minAvg := math.MaxInt64
	calAvg := 0
	start := 0

	for i := 0; i < len(A) && i+k <= len(A); i++ {
		j := i
		for ; j < i+k; j++ {
			calAvg += A[j]
		}
		calAvg /= k
		if calAvg < minAvg {
			minAvg = calAvg
			start = i
		}
		calAvg = 0

	}
	return A[start : start+k]
}

