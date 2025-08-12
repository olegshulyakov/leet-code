package main

import (
	"reflect"
	"testing"
)

func findDifference(nums1 []int, nums2 []int) [][]int {
	// Create maps
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)

	// Populate sets
	for _, num := range nums1 {
		set1[num] = true
	}
	for _, num := range nums2 {
		set2[num] = true
	}

	// Remove common elements
	for _, num := range nums2 {
		delete(set1, num)
	}
	for _, num := range nums1 {
		delete(set2, num)
	}

	// Convert map keys to slices
	result1 := make([]int, 0, len(set1))
	for num := range set1 {
		result1 = append(result1, num)
	}

	result2 := make([]int, 0, len(set2))
	for num := range set2 {
		result2 = append(result2, num)
	}

	return [][]int{result1, result2}
}

func TestFindDifference(t *testing.T) {
	testCases := []struct {
		nums1 []int
		nums2 []int
		want  [][]int
	}{
		{[]int{1, 2, 3}, []int{2, 4, 6}, [][]int{{1, 3}, {4, 6}}},
		{[]int{1, 2, 3, 3}, []int{1, 1, 2, 2}, [][]int{{3}, {}}},
	}

	for _, tc := range testCases {
		t.Run("2215. Find the Difference of Two Arrays", func(t *testing.T) {
			out := findDifference(tc.nums1, tc.nums2)

			if len(out) != len(tc.want) {
				t.Errorf("len(findDifference(%v, %v) = %v, want: %v", tc.nums1, tc.nums2, out, tc.want)
			}
			for i := range out {
				if !equalUnordered(out[i], tc.want[i]) {
					t.Errorf("findDifference()[%v] = %v, want: %v", i+1, out[i], tc.want[i])
				}
			}
		})
	}
}

// Compare slices regardless of order.
func equalUnordered(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	mapA := make(map[int]int)
	mapB := make(map[int]int)

	for _, v := range a {
		mapA[v]++
	}
	for _, v := range b {
		mapB[v]++
	}

	return reflect.DeepEqual(mapA, mapB)
}
