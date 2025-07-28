package main

import (
	"testing"
)

/*

Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.



Example 1:

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
Example 2:

Input: nums = [0]
Output: [0]


Constraints:

1 <= nums.length <= 10^4
-231 <= nums[i] <= 231 - 1


Follow up: Could you minimize the total number of operations done?

*/

// moveZeroes moves all zeroes to the end of the array while maintaining
// the relative order of the non-zero elements.
//
// The algorithm uses a two-pointer approach:
// 1. One pointer (i) tracks the position where the next non-zero element should go
// 2. Another pointer (counter) iterates through the array
// 3. When a non-zero element is found, it's moved to position i
// 4. After processing all elements, fill remaining positions with zeros
//
// Time Complexity: O(n), Space Complexity: O(1)
//
// Parameters:
//
//	nums []int - Slice of integers to modify in-place
func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}

	// Two-pointer approach
	left := 0
	for right := range nums {
		if nums[right] != 0 {
			nums[left] = nums[right]
			left++
		}
	}

	// Fill remaining with zeros
	for left < len(nums) {
		nums[left] = 0
		left++
	}
}

func TestMoveZeroes(t *testing.T) {
	testCases := []struct {
		nums []int
		want []int
	}{
		{nums: []int{0, 1, 0, 3, 12}, want: []int{1, 3, 12, 0, 0}},
		{nums: []int{0}, want: []int{0}},
	}
	for _, tc := range testCases {
		t.Run("283. Move Zeroes", func(t *testing.T) {
			out := make([]int, len(tc.nums))
			copy(out, tc.nums)

			moveZeroes(out)

			if len(out) != len(tc.want) {
				t.Errorf("len(moveZeroes() = %v, want: %v", len(out), len(tc.want))
			}
			for i := range out {
				if out[i] != tc.want[i] {
					t.Errorf("moveZeroes()[%v] = %v, want: %v", i+1, out[i], tc.want[i])
				}
			}
		})
	}
}
