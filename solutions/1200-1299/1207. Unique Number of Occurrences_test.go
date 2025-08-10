package main

import "testing"

/*

Given an array of integers arr, return true if the number of occurrences of each value in the array is unique or false otherwise.



Example 1:

Input: arr = [1,2,2,1,1,3]
Output: true
Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1. No two values have the same number of occurrences.
Example 2:

Input: arr = [1,2]
Output: false
Example 3:

Input: arr = [-3,0,1,-3,1,1,1,-3,10,0]
Output: true


Constraints:

1 <= arr.length <= 1000
-1000 <= arr[i] <= 1000

*/

func uniqueOccurrences(arr []int) bool {
	counts := make(map[int]int)

	for _, num := range arr {
		counts[num]++
	}

	occurrences := make(map[int]bool)
	for _, count := range counts {
		if occurrences[count] {
			return false
		}
		occurrences[count] = true
	}

	return true
}

func TestUniqueOccurrences(t *testing.T) {
	const name = "1207. Unique Number of Occurrences"
	testCases := []struct {
		arr  []int
		want bool
	}{
		{[]int{1, 2, 2, 1, 1, 3}, true},
		{[]int{1, 2}, false},
		{[]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, true},
	}
	for _, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			out := uniqueOccurrences(tc.arr)
			if out != tc.want {
				t.Errorf("uniqueOccurrences(%v) = %v, want: %v", tc.arr, out, tc.want)
			}
		})
	}
}
