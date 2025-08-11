package main

import (
	"testing"
)

/*

We are given an array asteroids of integers representing asteroids in a row.

For each asteroid, the absolute value represents its size, and the sign represents its direction (positive meaning right, negative meaning left). Each asteroid moves at the same speed.

Find out the state of the asteroids after all collisions. If two asteroids meet, the smaller one will explode. If both are the same size, both will explode. Two asteroids moving in the same direction will never meet.


Example 1:

Input: asteroids = [5,10,-5]
Output: [5,10]
Explanation: The 10 and -5 collide resulting in 10. The 5 and 10 never collide.


Example 2:

Input: asteroids = [8,-8]
Output: []
Explanation: The 8 and -8 collide exploding each other.


Example 3:

Input: asteroids = [10,2,-5]
Output: [10]
Explanation: The 2 and -5 collide resulting in -5. The 10 and -5 collide resulting in 10.


Constraints:

2 <= asteroids.length <= 10^4
-1000 <= asteroids[i] <= 1000
asteroids[i] != 0

*/

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)

	for _, asteroid := range asteroids {
		if len(stack) == 0 {
			stack = append(stack, asteroid)
			continue
		}

		if asteroid >= 0 {
			stack = append(stack, asteroid)
			continue
		}

		// Explode asteroids in stack
		idxLast := len(stack) - 1
		last := stack[idxLast]
		for idxLast > 0 && last > 0 && last < -asteroid {
			idxLast--
			last = stack[idxLast]
		}
		stack = stack[:idxLast]

		switch {
		case last <= 0:
			stack = append(stack, last)
			stack = append(stack, asteroid)
		case last > -asteroid:
			stack = append(stack, last)
		case last < -asteroid:
			stack = append(stack, asteroid)
		}
	}

	return stack
}

func TestAsteroidCollision(t *testing.T) {
	testCases := []struct {
		asteroids []int
		want      []int
	}{
		{[]int{}, []int{}},
		{[]int{5, 10, -5}, []int{5, 10}},
		{[]int{8, -8}, []int{}},
		{[]int{10, 2, -5}, []int{10}},
		{[]int{-2, -1, 1, 2}, []int{-2, -1, 1, 2}},
		{[]int{-2, -2, 1, -2}, []int{-2, -2, -2}},
		{[]int{-2, -2, 1, -1}, []int{-2, -2}},
		{[]int{-2, 2, 1, -2}, []int{-2}},
		{[]int{1, -2, -2, -2}, []int{-2, -2, -2}},
	}
	for _, tc := range testCases {
		t.Run("735. Asteroid Collision", func(t *testing.T) {
			t.Logf("asteroidCollision(%v)", tc.asteroids)
			out := asteroidCollision(tc.asteroids)
			if len(out) != len(tc.want) {
				t.Errorf("len(out) = %v, want: %v", out, tc.want)
			}
			for i := range out {
				if out[i] != tc.want[i] {
					t.Errorf("asteroidCollision()[%v] = %v, want: %v", i+1, out[i], tc.want[i])
				}
			}
		})
	}
}
