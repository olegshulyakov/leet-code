package main

import (
	"fmt"
	"strings"
	"testing"
)

// fun takes a slice of strings and returns the same slice unchanged.
// This is a simple identity function that demonstrates basic slice handling.
//
// Parameters:
//
//	s []string - Input slice of strings to be returned unchanged
//
// Returns:
//
//	[]string - The same slice that was passed in as input
func fun(s []string) []string {
	return s
}

// Test validates the functionality of the fun function using table-driven tests.
// It defines test cases and iterates through them to verify correct behavior.
func Test(t *testing.T) {
	testcases := []struct {
		in, want []string
	}{
		{[]string{}, []string{}},
	}
	for _, tc := range testcases {
		name := fmt.Sprintf("Test: %s", strings.Join(tc.in, ", "))
		t.Run(name, func(t *testing.T) {
			t.Logf("fun: in=%v", tc.in)
			out := fun(tc.in)
			if len(out) != len(tc.want) {
				t.Errorf("len(fun() = %v, want: %v", len(out), len(tc.want))
			}
			for i := range out {
				if out[i] != tc.want[i] {
					t.Errorf("fun()[%v] = %v, want: %v", i+1, out[i], tc.want[i])
				}
			}
		})
	}
}
