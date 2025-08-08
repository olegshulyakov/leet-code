package main

import "testing"

/*

You are given an integer array nums consisting of n elements, and an integer k.

Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. Any answer with a calculation error less than 10-5 will be accepted.



Example 1:

Input: nums = [1,12,-5,-6,50,3], k = 4
Output: 12.75000
Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75
Example 2:

Input: nums = [5], k = 1
Output: 5.00000


Constraints:

n == nums.length
1 <= k <= n <= 10^5
-104 <= nums[i] <= 10^4

*/

func findMaxAverage(nums []int, k int) float64 {
	if len(nums) < k {
		return 0
	}

	sum := 0
	temp := 0

	for i := range len(nums) {
		if i < k {
			temp += nums[i]
			sum = temp
		} else {
			temp += nums[i] - nums[i-k]
			if temp > sum {
				sum = temp
			}
		}
	}

	return float64(sum) / float64(k)
}

func TestFindMaxAverage(t *testing.T) {
	const name = "643. Maximum Average Subarray I"
	testCases := []struct {
		nums []int
		k    int
		want float64
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4, 12.75},
		{[]int{5}, 1, 5},
		{[]int{0, 1, 1, 3, 3}, 4, 2},
		{[]int{0, 4, 0, 3, 2}, 1, 4},
	}

	for _, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			out := findMaxAverage(tc.nums, tc.k)
			if out != tc.want {
				t.Errorf("findMaxAverage(%v, %d) = %v; expected %v", tc.nums, tc.k, out, tc.want)
			}
		})
	}
}
