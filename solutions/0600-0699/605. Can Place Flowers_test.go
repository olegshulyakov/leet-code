package arrayorstring_test

import (
	"testing"
)

/*
You have a long flowerbed in which some of the plots are planted, and some are not. However, flowers cannot be planted in adjacent plots.

Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n, return true if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule and false otherwise.

Example 1:
Input: flowerbed = [1,0,0,0,1], n = 1
Output: true

Example 2:
Input: flowerbed = [1,0,0,0,1], n = 2
Output: false


Constraints:
1 <= flowerbed.length <= 2 * 10^4
flowerbed[i] is 0 or 1.
There are no two adjacent flowers in flowerbed.
0 <= n <= flowerbed.length
*/

// canPlaceFlowers determines if 'n' flowers can be planted in the flowerbed
// without violating the no-adjacent-flowers rule.
// Flowers cannot be planted in adjacent plots.
//
// Parameters:
//
//	flowerbed []int - Array representing flowerbed where 0 = empty, 1 = occupied
//	n int - Number of flowers to plant
//
// Returns:
//
//	bool - True if 'n' flowers can be planted, false otherwise
func canPlaceFlowers(flowerbed []int, n int) bool {
	// If no flowers need to be planted, return true immediately
	if n == 0 {
		return true
	}

	// Iterate through each plot in the flowerbed
	for i := 0; i < len(flowerbed); i++ {
		// Check if current plot or adjacent plots are occupied
		// - flowerbed[i-1] == 1: left adjacent plot is occupied
		// - flowerbed[i] == 1: current plot is occupied
		// - flowerbed[i+1] == 1: right adjacent plot is occupied
		if (i > 0 && flowerbed[i-1] == 1) ||
			flowerbed[i] == 1 ||
			(i < len(flowerbed)-1 && flowerbed[i+1] == 1) {
			continue
		}

		// Found a valid spot to plant a flower
		i++ // Skip next plot (no adjacent planting allowed)
		n-- // Decrement remaining flowers to plant

		// If all flowers are planted, break early
		if n == 0 {
			break
		}
	}

	// Return true if all required flowers were planted
	return n == 0
}

func TestCanPlaceFlowers(t *testing.T) {
	testCases := []struct {
		in   []int
		n    int
		want bool
	}{
		{in: []int{1, 0, 0, 0, 1}, n: 1, want: true},
		{in: []int{1, 0, 0, 0, 1}, n: 2, want: false},
	}
	for _, tc := range testCases {
		t.Run("605. Can Place Flowers", func(t *testing.T) {
			out := canPlaceFlowers(tc.in, tc.n)
			if out != tc.want {
				t.Errorf("canPlaceFlowers(%v, %d) = %v, want: %v", tc.in, tc.n, out, tc.want)
			}
		})
	}
}
