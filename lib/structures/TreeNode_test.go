package structures_test

import (
	"reflect"
	"testing"

	"github.com/olegshulyakov/leet-code.go/lib/structures"
)

// TestTreeNode tests the TreeNode functionality.
func TestTreeNode(t *testing.T) {
	// Test creating a binary tree from an array of numbers
	t.Run("should create a binary tree from an array of numbers", func(t *testing.T) {
		in := []int{1, 2, 0, 3, 4}
		root := structures.NewTree(in)
		out := root.ToArray()

		if !reflect.DeepEqual(out, in) {
			t.Errorf("want %v, got %v", in, out)
		}
	})

	// Test creating a binary tree with no nodes
	t.Run("should create a binary tree with no nodes", func(t *testing.T) {
		in := []int{}
		root := structures.NewTree(in)
		out := root.ToArray()

		if !reflect.DeepEqual(out, in) {
			t.Errorf("want %v, got %v", in, out)
		}
	})

	// Skip test: should not create a binary tree with duplicate values
	t.Run("should not create a binary tree with duplicate values", func(t *testing.T) {
		t.Skip("Skipping duplicate values test")

		in := []int{1, 2, 0, 3, 4}
		modifiedNums := []int{1, 2, 2, 3, 4}
		root := structures.NewTree(modifiedNums)
		out := root.ToArray()

		if !reflect.DeepEqual(out, in) {
			t.Errorf("want %v, got %v", in, out)
		}
	})

	// Test creating a binary tree with duplicate values
	t.Run("should create a binary tree with duplicate values", func(t *testing.T) {
		in := []int{1, 2, 2, 3, 4}
		root := structures.NewTree(in)
		out := root.ToArray()

		if !reflect.DeepEqual(out, in) {
			t.Errorf("want %v, got %v", in, out)
		}
	})

	// Test creating a binary tree with only one node
	t.Run("should create a binary tree with only one node", func(t *testing.T) {
		in := []int{1}
		root := structures.NewTree(in)
		out := root.ToArray()

		if !reflect.DeepEqual(out, in) {
			t.Errorf("want %v, got %v", in, out)
		}
	})

	// Test that converting to array from nil node does not panic
	t.Run("should not throw an error when converting to array from null node", func(t *testing.T) {
		// This should not panic
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
	})
}
