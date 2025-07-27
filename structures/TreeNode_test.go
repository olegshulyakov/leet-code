package structures_test

import (
	"reflect"
	"testing"

	"github.com/olegshulyakov/leet-code.go/structures"
)

func TestTreeNode(t *testing.T) {
	testcases := []struct {
		in, want []int
	}{
		{[]int{1, 2, 0, 3, 4}, []int{1, 2, 0, 3, 4}},
		{[]int{}, []int{}},
		{[]int{1, 2, 2, 3, 4}, []int{1, 2, 2, 3, 4}},
		{[]int{1}, []int{1}},
	}

	for _, tc := range testcases {
		root := structures.NewTree(tc.in)
		out := root.ToArray()

		// Verify the length of output matches expected length
		if len(out) != len(tc.want) {
			t.Errorf("Length mismatch - want: %v, got: %v", len(tc.want), len(out))
		}

		// Compare each element of the output with expected values
		for i := range out {
			if out[i] != tc.want[i] {
				t.Errorf("Element mismatch at index [%v] - want: %q, got: %q",
					i+1, tc.want[i], out[i])
			}
		}
	}
}

func TestTreeNodeWithNil(t *testing.T) {
	// Test that converting to array from nil node does not panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("TreeNode.ToArray(nil) panicked: %v", r)
		}
	}()

	var root *structures.TreeNode
	out := root.ToArray()
	want := []int{}

	if !reflect.DeepEqual(out, want) {
		t.Errorf("want %v, got %v", want, out)
	}
}
