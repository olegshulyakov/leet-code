package main

import (
	"testing"
)

func equalPairs(grid [][]int) int {
	var rows [][]int
	var cols [][]int
	for i := range grid {
		var row []int
		var column []int
		for j := range grid {
			row = append(row, grid[i][j])
			column = append(column, grid[j][i])
		}
		rows = append(rows, row)
		cols = append(cols, column)
	}

	pairCnt := 0
	for _, row := range rows {
		for _, col := range cols {
			if len(row) != len(col) {
				continue
			}
			var isEqual = true
			for i := range row {
				if row[i] != col[i] {
					isEqual = false
					break
				}
			}
			if !isEqual {
				continue
			}
			pairCnt++
		}
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
