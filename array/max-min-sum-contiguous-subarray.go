package array

import (
	"math"
)

// Maximum Sum SubArray
// (Largest Sum Contigous SubArray)
func MaximumSumSubArray(A []int) int {
	var localSum, globalSum int = A[0], A[0]
	for i := 1; i < len(A); i++ {
		localSum = int(math.Max(float64(A[i]), float64(A[i]+localSum)))
		if localSum >= globalSum {
			globalSum = localSum
		}
	}
	return globalSum
}

// Minimum Sum SubArray
// (Smallest Sum Contigous SubArray)
func MinimumSumSubArray(A []int) int {
	var localSum, globalSum int = A[0], A[0]
	for i := 1; i < len(A); i++ {
		localSum = int(math.Min(float64(A[i]), float64(A[i] + localSum)))
		if localSum < globalSum {
			globalSum = localSum
		}
	}
	return  globalSum
}
