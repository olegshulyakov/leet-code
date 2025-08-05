package main

import "testing"

/*

ou are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.



Example 1:


Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
Example 2:

Input: height = [1,1]
Output: 1


Constraints:

n == height.length
2 <= n <= 10^5
0 <= height[i] <= 10^4

*/

// maxArea finds the maximum area of water that can be stored between two lines.
// Uses the two-pointer technique to efficiently find the optimal solution.
//
// The algorithm:
// 1. Start with pointers at both ends of the array
// 2. Calculate area formed by the two pointers
// 3. Move the pointer pointing to the shorter line inward
// 4. Continue until pointers meet
//
// This works because moving the taller line's pointer can only decrease the area,
// while moving the shorter line's pointer might find a taller line that results
// in a larger area despite the decreased width.
//
// Parameters:
//
//	height []int - Array representing heights of vertical lines
//
// Returns:
//
//	int - Maximum area of water that can be stored
func maxArea(height []int) int {
	// Handle edge case: need at least 2 lines to form a container
	if len(height) < 2 {
		return 0
	}

	areaSize := 0
	left := 0
	right := len(height) - 1

	// Two-pointer approach
	for left < right {
		current := min(height[left], height[right]) * (right - left)
		areaSize = max(areaSize, current)

		// Move the pointer pointing to the shorter line
		// This is the key insight: moving the taller line can only decrease area
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return areaSize
}

func TestMaxArea(t *testing.T) {
	const name = "11. Container With Most Water"
	// Array of 10000 ones
	ones := make([]int, 10000)
	for i := range ones {
		ones[i] = 1
	}
	testCases := []struct {
		height []int
		want   int
	}{
		{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, want: 49},
		{height: []int{1, 1}, want: 1},
		{height: []int{0}, want: 0},
		{height: []int{1, 2, 1}, want: 2},
		{height: []int{1, 2, 4, 3}, want: 4},
		{height: []int{2, 3, 4, 5, 18, 17, 6}, want: 17},
		{height: []int{0, 14, 6, 2, 10, 9, 4, 1, 10, 3}, want: 70},
		{height: ones, want: 9999},
	}
	for _, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			out := maxArea(tc.height)
			if out != tc.want {
				t.Errorf("maxArea(%v) = %v, want: %v", tc.height, out, tc.want)
			}
		})
	}
}
