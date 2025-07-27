package main

import (
	"testing"
)

/*
Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k]. If no such indices exists, return false.


Example 1:
Input: nums = [1,2,3,4,5]
Output: true
Explanation: Any triplet where i < j < k is valid.

Example 2:
Input: nums = [5,4,3,2,1]
Output: false
Explanation: No triplet exists.

Example 3:
Input: nums = [2,1,5,0,4,6]
Output: true
Explanation: The triplet (3, 4, 5) is valid because nums[3] == 0 < nums[4] == 4 < nums[5] == 6.


Constraints:

1 <= nums.length <= 5 * 10^5
-231 <= nums[i] <= 231 - 1


Follow up: Could you implement a solution that runs in O(n) time complexity and O(1) space complexity?
*/

// increasingTriplet checks if there exists an increasing subsequence of length 3 in the array.
// It returns true if there are indices i < j < k such that nums[i] < nums[j] < nums[k].
// Uses an optimized approach with O(n) time complexity and O(1) space complexity.
//
// Parameters:
//
//	nums []int - Input array of integers
//
// Returns:
//
//	bool - True if an increasing triplet subsequence exists, false otherwise
func increasingTriplet(nums []int) bool {
	// Need at least 3 elements for a triplet
	if len(nums) < 3 {
		return false
	}

	// Track the smallest number seen so far
	smallestNumber := nums[0]

	// Track the smallest second number in potential increasing pairs
	smallestSecond := int(^uint(0) >> 1) // Max int value

	// Iterate through the array starting from index 1
	for i := 1; i < len(nums); i++ {
		numberToCheck := nums[i]

		// If current number is greater than smallest second, we found a triplet
		if numberToCheck > smallestSecond {
			return true
		}

		// If current number is greater than smallest but less than smallestSecond,
		// update smallestSecond (this forms a new potential pair)
		if numberToCheck > smallestNumber && numberToCheck < smallestSecond {
			smallestSecond = numberToCheck
		}

		// If current number is smaller than smallest, update smallest
		if numberToCheck < smallestNumber {
			smallestNumber = numberToCheck
		}
	}

	return false
}

func TestIncreasingTriplet(t *testing.T) {
	testCases := []struct {
		in   []int
		want bool
	}{
		{in: []int{1, 2, 3, 4, 5}, want: true},
		{in: []int{5, 4, 3, 2, 1}, want: false},
		{in: []int{2, 1, 5, 0, 4, 6}, want: true},
		{in: []int{1, 2, 3}, want: true},
		{in: []int{1, 2}, want: false},
		{in: []int{1, 3, 2, 4}, want: true},
		{in: []int{20, 100, 10, 12, 5, 13}, want: true},
		{in: []int{1, 1, -2, 6}, want: false},
		{in: make([]int, 50), want: false}, // Array of 50 zeros
	}
	for _, tc := range testCases {
		t.Run("334. Increasing Triplet Subsequence", func(t *testing.T) {
			out := increasingTriplet(tc.in)
			if out != tc.want {
				t.Errorf("increasingTriplet(%v) = %v, want: %v", tc.in, out, tc.want)
			}
		})
	}
}
