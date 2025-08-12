package main

import "testing"

func minReorder(n int, connections [][]int) int {
	type edge struct {
		city  int
		isOut bool
	}

	roads := make(map[int][]edge, n)
	for _, conn := range connections {
		a, b := conn[0], conn[1]
		roads[a] = append(roads[a], edge{city: b, isOut: true})
		roads[b] = append(roads[b], edge{city: a, isOut: false})
	}

	visited := make([]bool, n)
	flips := 0
	var goToCity func(int)
	goToCity = func(node int) {
		visited[node] = true
		for _, e := range roads[node] {
			if visited[e.city] {
				continue
			}
			if e.isOut {
				flips++
			}
			goToCity(e.city)
		}
	}

	goToCity(0)
	return flips
}

func TestMinReorder(t *testing.T) {
	testCases := []struct {
		n           int
		connections [][]int
		want        int
	}{
		{6, [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}, 3},
		{5, [][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}}, 2},
		{3, [][]int{{1, 0}, {2, 0}}, 0},
		{6, [][]int{{4, 5}, {0, 1}, {1, 3}, {2, 3}, {4, 0}}, 3},
	}

	for _, tc := range testCases {
		t.Run("1466. Reorder Routes to Make All Paths Lead to the City Zero", func(t *testing.T) {
			out := minReorder(tc.n, tc.connections)
			if out != tc.want {
				t.Fatalf("minReorder(%d, %v) = %d; want %d", tc.n, tc.connections, out, tc.want)
			}
		})
	}
}
