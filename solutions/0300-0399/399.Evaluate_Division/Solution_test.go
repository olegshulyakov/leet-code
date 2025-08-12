package main

import (
	"testing"
)

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	if len(equations) == 0 || len(values) == 0 || len(queries) == 0 {
		return []float64{}
	}

	evals := make(map[string]map[string]float64)
	addToMap := func(dividend, divisor string, value float64) {
		if _, ok := evals[dividend]; !ok {
			evals[dividend] = make(map[string]float64)
		}
		evals[dividend][divisor] = value
	}

	for i, eq := range equations {
		if len(eq) != 2 {
			continue
		}
		a, b := eq[0], eq[1]
		if a == b {
			continue
		}
		val := values[i]
		addToMap(a, b, val)
		addToMap(b, a, 1.0/val)
	}

	answers := make([]float64, len(queries))
	for idx, query := range queries {
		if len(query) != 2 {
			answers[idx] = -1.0
			continue
		}

		a, b := query[0], query[1]
		if _, ok1 := evals[a]; !ok1 {
			answers[idx] = -1.0
		} else if _, ok2 := evals[b]; !ok2 {
			answers[idx] = -1.0
		} else if a == b {
			answers[idx] = 1.0
		} else if val, ok3 := evals[a][b]; ok3 {
			answers[idx] = val
		} else {
			visited := make(map[string]bool)
			answers[idx] = dfs(evals, a, b, visited)
		}
	}
	return answers
}

func dfs(evals map[string]map[string]float64, dividend, divisor string, visited map[string]bool) float64 {
	visited[dividend+"|"+divisor] = true

	if memorized, ok1 := evals[dividend]; ok1 {
		if val, ok2 := memorized[divisor]; ok2 {
			return val
		}

		for next, val := range memorized {
			if visited[next+"|"+divisor] {
				continue
			}
			if res := dfs(evals, next, divisor, visited); res != -1.0 {
				return val * res
			}
		}
	}
	return -1.0
}

func TestCalcEquation(t *testing.T) {
	testCases := []struct {
		equations [][]string
		values    []float64
		queries   [][]string
		want      []float64
	}{
		{
			[][]string{},
			[]float64{},
			[][]string{},
			[]float64{},
		},
		{
			[][]string{{"a", "b"}, {"b", "c"}},
			[]float64{2.0, 3.0},
			[][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
			[]float64{6.0, 0.5, -1.0, 1.0, -1.0},
		},
		{
			[][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}},
			[]float64{1.5, 2.5, 5.0},
			[][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}},
			[]float64{3.75, 0.4, 5.0, 0.2},
		},
		{
			[][]string{{"a", "b"}},
			[]float64{0.5},
			[][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}},
			[]float64{0.5, 2.0, -1.0, -1.0},
		},
		{
			[][]string{{"x1", "x2"}, {"x2", "x3"}, {"x3", "x4"}, {"x4", "x5"}},
			[]float64{3.0, 4.0, 5.0, 6.0},
			[][]string{{"x1", "x5"}, {"x5", "x2"}, {"x2", "x4"}, {"x2", "x2"}, {"x2", "x9"}, {"x9", "x9"}},
			[]float64{360.0, 0.008333333333333333, 20.0, 1.0, -1.0, -1.0},
		},
		{
			[][]string{{"a", "b"}, {"c", "d"}},
			[]float64{1.0, 1.0},
			[][]string{{"a", "c"}, {"b", "d"}, {"b", "a"}, {"d", "c"}},
			[]float64{-1.0, -1.0, 1.0, 1.0},
		},
	}
	for _, tc := range testCases {
		t.Run("399. Evaluate Division", func(t *testing.T) {
			out := calcEquation(tc.equations, tc.values, tc.queries)
			if !compare(out, tc.want) {
				t.Errorf("calcEquation(%v, %v, %v) = %v, want %v", tc.equations, tc.values, tc.queries, out, tc.want)
			}
		})
	}
}

func compare(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	eps := 1e-5
	for i := range a {
		if (a[i]-b[i]) > eps || (b[i]-a[i]) > eps {
			return false
		}
	}
	return true
}
