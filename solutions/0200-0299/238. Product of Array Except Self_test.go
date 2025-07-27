package arrayorstring_test

import (
	"testing"
)

/*

Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
You must write an algorithm that runs in O(n) time and without using the division operation.


Example 1:
Input: nums = [1,2,3,4]
Output: [24,12,8,6]

Example 2:
Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]


Constraints:

2 <= nums.length <= 10^5
-30 <= nums[i] <= 30
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.


Follow up: Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)

*/

// productExceptSelf calculates the product of all elements in the array except the current element.
// For each index i, the result[i] contains the product of all elements except nums[i].
// Handles special cases with zeros efficiently.
//
// Parameters:
//
//	nums []int - Input array of integers
//
// Returns:
//
//	[]int - Array where each element is the product of all other elements
func productExceptSelf(nums []int) []int {
	// Find indices of all zero elements
	zeros := []int{}
	for i, element := range nums {
		if element == 0 {
			zeros = append(zeros, i)
		}
	}

	// If there are 2 or more zeros, all products will be zero
	if len(zeros) >= 2 {
		res := make([]int, len(nums))
		return res
	}

	// Calculate total product of all non-zero elements
	total := 1
	for _, currentValue := range nums {
		if currentValue != 0 {
			total *= currentValue
		}
	}

	// If there is exactly one zero, all elements are zero except at zero's position
	if len(zeros) == 1 {
		res := make([]int, len(nums))
		res[zeros[0]] = total
		return res
	}

	// If there are no zeros, divide total by each element
	res := make([]int, len(nums))
	for i, element := range nums {
		if element != 0 {
			res[i] = total / element
		} else {
			res[i] = total
		}
	}

	return res
}

func TestProductExceptSelf(t *testing.T) {
	testCases := []struct{ in, want []int }{
		{in: []int{1, 2, 3, 4}, want: []int{24, 12, 8, 6}},
		{in: []int{-1, 1, 0, -3, 3}, want: []int{0, 0, 9, 0, 0}},
		{in: []int{-1, 1, 0, 0, 3}, want: []int{0, 0, 0, 0, 0}},
	}
	for _, tc := range testCases {
		t.Run("238. Product of Array Except Self", func(t *testing.T) {
			out := productExceptSelf(tc.in)

			if len(out) != len(tc.want) {
				t.Errorf("len(productExceptSelf() = %v, want: %v", len(out), len(tc.want))
			}
			for i := range out {
				if out[i] != tc.want[i] {
					t.Errorf("productExceptSelf()[%v] = %q, want: %q", i+1, out[i], tc.want[i])
				}
			}
		})
	}
}
