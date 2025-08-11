package main

import "testing"

/*

Given a binary array nums, you should delete one element from it.

Return the size of the longest non-empty subarray containing only 1's in the resulting array. Return 0 if there is no such subarray.



Example 1:

Input: nums = [1,1,0,1]
Output: 3
Explanation: After deleting the number in position 2, [1,1,1] contains 3 numbers with value of 1's.
Example 2:

Input: nums = [0,1,1,1,0,1,1,0,1]
Output: 5
Explanation: After deleting the number in position 4, [0,1,1,1,1,1,0,1] longest subarray with value of 1's is [1,1,1,1,1].
Example 3:

Input: nums = [1,1,1]
Output: 2
Explanation: You must delete one element.


Constraints:

1 <= nums.length <= 10^5
nums[i] is either 0 or 1.

*/

func longestSubarray(nums []int) int {
	left := 0
	maximum := 0
	hasZero := false

	for i := range nums {
		if nums[i] == 0 {
			if hasZero {
				for nums[left] != 0 {
					left++
				}
				left++
			} else {
				hasZero = true
			}
		}

		diff := i - left + 1
		if maximum < diff {
			maximum = diff
		}
	}

	return maximum - 1
}

func TestLongestSubarray(t *testing.T) {
	const name = "1493. Longest Subarray of ones After Deleting One Element"
	testCases := []struct {
		nums []int
		want int
	}{
		{[]int{1, 1, 0, 1}, 3},
		{[]int{0, 1, 1, 1, 0, 1, 1, 0, 1}, 5},
		{[]int{1, 1, 1}, 2},
	}
	for _, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Logf("longestSubarray(%v)", tc.nums)
			out := longestSubarray(tc.nums)
			if out != tc.want {
				t.Errorf("longestSubarray(%v) = %v, want: %v", tc.nums, out, tc.want)
			}
		})
	}
}
