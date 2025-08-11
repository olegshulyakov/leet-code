package main

import (
	"testing"
)

/*

You are given an integer array nums and an integer k.

In one operation, you can pick two numbers from the array whose sum equals k and remove them from the array.

Return the maximum number of operations you can perform on the array.



Example 1:

Input: nums = [1,2,3,4], k = 5
Output: 2
Explanation: Starting with nums = [1,2,3,4]:
- Remove numbers 1 and 4, then nums = [2,3]
- Remove numbers 2 and 3, then nums = []
There are no more pairs that sum up to 5, hence a total of 2 operations.
Example 2:

Input: nums = [3,1,3,4,3], k = 6
Output: 1
Explanation: Starting with nums = [3,1,3,4,3]:
- Remove the first two 3's, then nums = [1,4,3]
There are no more pairs that sum up to 6, hence a total of 1 operation.


Constraints:

1 <= nums.length <= 10^5
1 <= nums[i] <= 10^9
1 <= k <= 10^9

*/

// maxOperations finds the maximum number of operations where two numbers
// from the array sum to k, and each number can only be used once.
//
// The algorithm uses a hash map approach:
// 1. For each number, calculate its complement (k - number)
// 2. If complement exists in map, form a pair and increment operations
// 3. Otherwise, add the current number to the map
// 4. Track counts of each number to handle duplicates properly
//
// Time Complexity: O(n) where n is the length of nums
// Space Complexity: O(n) for the hash map
//
// Parameters:
//
//	nums []int - Array of integers
//	k int - Target sum for pairs
//
// Returns:
//
//	int - Maximum number of k-sum pairs that can be formed
func maxOperations(nums []int, k int) int {
	// Handle edge case: need at least 2 elements to form a pair
	if len(nums) < 2 {
		return 0
	}

	ops := 0
	frequencyMap := make(map[int]int)
	for _, current := range nums {
		// Skip numbers greater than k (cannot form valid pairs)
		if current > k {
			continue
		}

		difference := k - current
		// Check if difference exists in map
		if counter, exists := frequencyMap[difference]; exists {
			// Form a pair
			if counter > 1 {
				// Decrease count of difference
				frequencyMap[difference] = counter - 1
			} else {
				// Remove difference from map
				delete(frequencyMap, difference)
			}
			ops++
		} else if counter, exists = frequencyMap[current]; exists {
			// Increase count of current number
			frequencyMap[current] = counter + 1
		} else {
			// Add current number to map
			frequencyMap[current] = 1
		}
	}

	return ops
}

func TestMaxOperations(t *testing.T) {
	largeArray := make([]int, 10000)
	for i := range largeArray {
		largeArray[i] = 5000000
	}

	testCases := []struct {
		nums []int
		k    int
		want int
	}{
		{nums: []int{1, 2, 3, 4}, k: 5, want: 2},
		{nums: []int{3, 1, 3, 4, 3}, k: 6, want: 1},
		{nums: []int{2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2}, k: 3, want: 4},
		{nums: []int{1}, k: 2, want: 0},
		{nums: []int{}, k: 5, want: 0},
		{nums: []int{1, 1, 1, 1}, k: 2, want: 2},
		{nums: largeArray, k: 10000000, want: 5000},
	}
	for _, tc := range testCases {
		t.Run("1679. Max Number of K-Sum Pairs", func(t *testing.T) {
			out := maxOperations(tc.nums, tc.k)

			if out != tc.want {
				t.Errorf("maxOperations(%v, %v) = %v, want: %v", tc.nums, tc.k, out, tc.want)
			}
		})
	}
}
