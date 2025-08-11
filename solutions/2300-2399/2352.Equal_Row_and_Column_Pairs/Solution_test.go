package main

import (
	"fmt"
	"testing"
)

func equalPairs(grid [][]int) int {
	rowMap := make(map[string]int)
	for _, row := range grid {
		rowStr := fmt.Sprint(row)
        rowMap[rowStr]++
	}

	pairCnt := 0
	for i := range grid {
		column := make([]int, len(grid))
		for j := range grid {
			column[j] = grid[j][i]
		}
		colStr := fmt.Sprint(column)
        pairCnt += rowMap[colStr]
	}

	return pairCnt
}

func TestEqualPairs(t *testing.T) {
	testCases := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}, 1},
		{[][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}, 3},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := equalPairs(tc.grid)
			if result != tc.want {
				t.Errorf("equalPairs(%v) = %d; expected %d", tc.grid, result, tc.want)
			}
		})
	}
}
