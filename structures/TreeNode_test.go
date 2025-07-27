package structures_test

import (
	"testing"

	"github.com/olegshulyakov/leet-code.go/structures"
)

func TestTreeNode(t *testing.T) {
	testcases := []struct {
		in, want []any
	}{
		{[]any{1, 2, 0, 3, 4}, []any{1, 2, 0, 3, 4}},
		{[]any{}, []any{}},
		{[]any{1, 2, 2, 3, 4}, []any{1, 2, 2, 3, 4}},
		{[]any{1}, []any{1}},
	}

	for _, tc := range testcases {
		out := structures.NewTree(tc.in).ToArray()

		// Verify the length of output matches expected length
		if len(out) != len(tc.want) {
			t.Errorf("Length mismatch - want: %v, got: %v", len(tc.want), len(out))
		}

		// Compare each element of the output with expected values
		for i := range out {
			if out[i] != tc.want[i] {
				t.Errorf("Element mismatch at index [%v] - want: %v, got: %v",
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

	if len(out) != len(want) {
		t.Errorf("len() = %v, want: %v", len(out), len(want))
	}
}
