package array

import (
	"math"
)

// Given an unsorted array arr[] and two numbers x and y,
//find the minimum distance between x and y in arr[].
//The array might also contain duplicates.
//You may assume that both x and y are different and present in arr[].
//
//Examples:
//Input: arr[] = {1, 2}, x = 1, y = 2
//Output: Minimum distance between 1 and 2 is 1.
//
//Input: arr[] = {3, 4, 5}, x = 3, y = 5
//Output: Minimum distance between 3 and 5 is 2.

func MinDistanceBetweenNumbers(A []int, x, y int) int {
	minCount := math.MaxInt64
	currentCount := 0
	found := false

	for i := 0; i < len(A); i++ {
		if A[i] == x {
			found = true
		}
		if found && A[i] != y {
			currentCount++
		}

		if A[i] == y {
			if minCount > currentCount {
				minCount = currentCount
				currentCount = 0
				found = false
			}
		}
	}
	return minCount
}
