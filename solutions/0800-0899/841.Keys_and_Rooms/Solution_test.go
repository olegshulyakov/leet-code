package main

import "testing"

func canVisitAllRooms(rooms [][]int) bool {
	if len(rooms) == 0 {
		return true
	}

	visited := make([]bool, len(rooms))
	queue := []int{0}
	for len(queue) > 0 {
		next := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if visited[next] {
			continue
		}
		visited[next] = true

		for _, key := range rooms[next] {
			if key >= 0 && key < len(rooms) && !visited[key] {
				queue = append(queue, key)
			}
		}
	}

	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true
}

func TestCanVisitAllRooms(t *testing.T) {
	testCases := []struct {
		rooms [][]int
		want  bool
	}{
		{[][]int{}, true},
		{[][]int{{1}, {2}, {3}, {}}, true},
		{[][]int{{1, 3}, {3, 0, 1}, {2}, {0}}, false},
		{[][]int{{2}, {}, {1}}, true},
		{[][]int{{1, 3}, {1, 4}, {2, 3, 4, 1}, {}, {4, 3, 2}}, true},
	}
	for _, tc := range testCases {
		t.Run("841. Keys and Rooms", func(t *testing.T) {
			out := canVisitAllRooms(tc.rooms)
			if out != tc.want {
				t.Errorf("canVisitAllRooms(%v) = %v; want %v", tc.rooms, out, tc.want)
			}
		})
	}
}
