package main

import (
	"testing"

	"github.com/olegshulyakov/leet-code.go/structures"
)

/*

You are given the root of a binary search tree (BST) and an integer val.

Find the node in the BST that the node's value equals val and return the subtree rooted with that node. If such a node does not exist, return null.


Example 1:

Input: root = [4,2,7,1,3], val = 2
Output: [2,1,3]


Example 2:

Input: root = [4,2,7,1,3], val = 5
Output: []


Constraints:

The number of nodes in the tree is in the range [1, 5000].
1 <= Node.val <= 10^7
root is a binary search tree.
1 <= val <= 10^7

*/

// searchBST searches for a node with the specified value in a binary search tree.
// It returns the subtree rooted at the found node, or nil if the value is not found.
// This implementation works for any binary tree (not just BSTs) but is typically
// used with binary search trees where it can be optimized.
//
// Parameters:
//
//	root *TreeNode - Root of the binary tree to search
//	val int - Value to search for
//
// Returns:
//
//	*TreeNode - Subtree rooted at the found node, or nil if not found
func searchBST(root *structures.TreeNode, val int) *structures.TreeNode {
	// Base case: empty tree or subtree
	if root == nil {
		return nil
	}

	// Found the target node
	if root.Val == val {
		return root
	}

	// Search in left subtree
	leftResult := searchBST(root.Left, val)
	if leftResult != nil {
		return leftResult
	}

	// Search in right subtree
	rightResult := searchBST(root.Right, val)
	return rightResult
}

func TestDeleteNode(t *testing.T) {
	testCases := []struct {
		nums []any
		val  int
		want []any
	}{
		{nums: []any{4, 2, 7, 1, 3}, val: 2, want: []any{2, 1, 3}},
		{nums: []any{4, 2, 7, 1, 3}, val: 5, want: []any{}},
	}
	for _, tc := range testCases {
		t.Run("700. Search in a Binary Search Tree", func(t *testing.T) {
			t.Logf("Search in a Binary Search Tree: nums=%d, key=%d", tc.nums, tc.val)
			root := structures.NewTree(tc.nums)

			newRoot := searchBST(root, tc.val)
			t.Logf("Search in a Binary Search Tree: root:\n%v, newRoot:\n%v", root.Print(), newRoot.Print())
			out := newRoot.ToArray()
			if len(out) != len(tc.want) {
				t.Errorf("len(out) = %v, want: %v", len(out), len(tc.want))
			}
			for i := range out {
				if out[i] != tc.want[i] {
					t.Errorf("deleteNode()[%v] = %v, want: %v", i+1, out[i], tc.want[i])
				}
			}
		})
	}
}
