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

// Input: arr[] = {3, 5, 4, 2, 6, 5, 6, 6, 5, 4, 8, 3}, x = 3, y = 6
//Output: Minimum distance between 3 and 6 is 4.
//
//Input: arr[] = {2, 5, 3, 5, 4, 4, 2, 3}, x = 3, y = 2
//Output: Minimum distance between 3 and 2 is 1.

// Notice that order is NOT important means (x,y) and (y,x) are the case
// Time complexity O(n)
func MinDistanceBetweenNumbers(A []int, x, y int) int {
	minCount := math.MaxInt64
	currentCount := 0
	found := false
	start := math.MinInt64
	for i := 0; i < len(A); i++ {
		if (A[i] == x || A[i] == y ) && found == false {
			found = true
			start = A[i]
		}
		if found == true  && A[i] != start{
			currentCount++
		}
		if (A[i] == x || A[i] == y) && A[i] != start {
			if minCount > currentCount {
				minCount = currentCount
			}
			currentCount = 0
			start = math.MinInt64
			found = false
		}
	}
	return minCount
}
