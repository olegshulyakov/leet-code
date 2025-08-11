package main

import "testing"

/*

Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array if you can flip at most k 0's.


Example 1:

Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
Output: 6
Explanation: [0,1,1,1,1,0]


Example 2:

Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
Output: 10
Explanation: [1,1,0,0,1,1,1,0,1,1]


Constraints:

1 <= nums.length <= 10^5
nums[i] is either 0 or 1.
0 <= k <= nums.length

*/

func longestOnes(nums []int, k int) int {
	left, maximum, zeros := 0, 0, 0

	for i := range nums {
		if nums[i] == 0 {
			if zeros == k {
				for nums[left] != 0 {
					left++
				}
				left++
			} else {
				zeros++
			}
		}

		diff := i - left + 1
		if maximum < diff {
			maximum = diff
		}
	}

	return maximum
}

func TestLongestOnes(t *testing.T) {
	const name = "1004. Max Consecutive Ones III"
	testCases := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 1, 1}, 2, 3},
		{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
		{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3, 10},
		{make([]int, 10), 2, 2},
		{make([]int, 10), 10, 10},
	}

	for _, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			out := longestOnes(tc.nums, tc.k)
			if out != tc.want {
				t.Errorf("longestOnes(%v, %d) = %d; expected %d", tc.nums, tc.k, out, tc.want)
			}
		})
	}
}
