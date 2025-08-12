package main

import (
	"testing"
)

func findCircleNum(isConnected [][]int) int {
	if len(isConnected) == 0 {
		return 0
	}

	isVisited := make([]bool, len(isConnected))
	var visitNearbyTowns func(city int)
	visitNearbyTowns = func(city int) {
		for neighbor := range isConnected {
			if isVisited[neighbor] || isConnected[city][neighbor] == 0 {
				continue
			}

			isVisited[neighbor] = true
			visitNearbyTowns(neighbor)
		}
	}

	provinces := 0
	for city := range isConnected {
		if isVisited[city] {
			continue
		}

		isVisited[city] = true
		visitNearbyTowns(city)

		provinces++
	}
	return provinces
}

func TestFindCircleNum(t *testing.T) {
	tests := []struct {
		isConnected [][]int
		want        int
	}{
		{isConnected: [][]int{}, want: 0},
		{isConnected: [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}, want: 2},
		{isConnected: [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}, want: 3},
		{isConnected: [][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}}, want: 1},
	}

	for _, tt := range tests {
		t.Run("547. Number of Provinces", func(t *testing.T) {
			out := findCircleNum(tt.isConnected)
			if out != tt.want {
				t.Errorf("findCircleNum(%v) = %d; expected %d", tt.isConnected, out, tt.want)
			}
		})
	}
}
