package array

// Given an array of integers, A of length N, find out the maximum sum sub-array of non negative numbers from A.
//
//The sub-array should be contiguous i.e., a sub-array created by choosing the second and fourth element and skipping the third element is invalid.
//
//Maximum sub-array is defined in terms of the sum of the elements in the sub-array.
//
//Find and return the required subarray.
//
//NOTE:
//
//    1. If there is a tie, then compare with segment's length and return segment which has maximum length.
//    2. If there is still a tie, then return the segment with minimum starting index.

// Input 1:
//    A = [1, 2, 5, -7, 2, 3]
//
//Output 1:
//    [1, 2, 5]
//
//Explanation 1:
//    The two sub-arrays are [1, 2, 5] [2, 3].
//    The answer is [1, 2, 5] as its sum is larger than [2, 3].
//
//Input 2:
//    A = [10, -1, 2, 3, -4, 100]
//
//Output 2:
//    [100]
//
//Explanation 2:
//    The three sub-arrays are [10], [2, 3], [100].
//    The answer is [100] as its sum is larger than the other two.
//Input 3:
//    A = [0, 0, -1, 0 ]
//Output 2:
//    [0, 0]

func MaxSet(A []int) []int {
	start := 0
	end := 0
	maxSum := 0
	sum := 0

	for i := 0; i < len(A); {
		j := i
		for ; j < len(A) && A[j] >= 0; {
			sum += A[i]
			j++
		}
		if sum > maxSum {
			maxSum = sum
			start = i
			end = j
		} else if sum == maxSum && end-start < j-i {
			start = i
			end = j
		}
		i = j + 1
	}
	return A[start:end]
}
