package main

import (
	"testing"

	"github.com/olegshulyakov/leet-code.go/structures"
)

/*

Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”


Example 1:

Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
Output: 3
Explanation: The LCA of nodes 5 and 1 is 3.


Example 2:

Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
Output: 5
Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant of itself according to the LCA definition.


Example 3:

Input: root = [1,2], p = 1, q = 2
Output: 1


Constraints:

The number of nodes in the tree is in the range [2, 10^5].
-10^9 <= Node.val <= 10^9
All Node.val are unique.
p != q
p and q will exist in the tree.

*/

// lowestCommonAncestor finds the lowest common ancestor (LCA) of two nodes in a binary tree.
// The LCA is the deepest node that has both p and q as descendants (where a node can be
// a descendant of itself).
//
// The algorithm uses a recursive approach:
// 1. If current node is nil or matches either p or q, return current node
// 2. Recursively search for p and q in left and right subtrees
// 3. If both left and right searches return non-nil, current node is LCA
// 4. If only one search returns non-nil, that result is the LCA
//
// Parameters:
//
//	root *TreeNode - Root of the binary tree
//	p *TreeNode - First target node
//	q *TreeNode - Second target node
//
// Returns:
//
//	*TreeNode - Lowest common ancestor of p and q
func lowestCommonAncestor(root, p, q *structures.TreeNode) *structures.TreeNode {
	// Base case: root is nil or matches either p or q
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}

	// Recursively search in left and right subtrees
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// If both left and right are non-nil, current node is LCA
	if left != nil && right != nil {
		return root
	}

	// Return the non-nil result (either left or right, or nil if both are nil)
	if left != nil {
		return left
	}
	return right
}

func TestLowestCommonAncestor(t *testing.T) {
	testcases := []struct {
		in   []any
		p    int
		q    int
		want int
	}{
		{in: []any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}, p: 5, q: 1, want: 3},
		{in: []any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}, p: 5, q: 4, want: 5},
		{in: []any{1, 2}, p: 1, q: 2, want: 1},
	}
	for _, tc := range testcases {
		t.Run("236. Lowest Common Ancestor of a Binary Tree", func(t *testing.T) {
			t.Logf("lowestCommonAncestor: in=%v", tc.in)
			out := lowestCommonAncestor(structures.NewTree(tc.in), &structures.TreeNode{Val: tc.p}, &structures.TreeNode{Val: tc.q})
			if out.Val != tc.want {
				t.Errorf("lowestCommonAncestor(%v,%v,%v) = %v, want: %v", tc.in, tc.p, tc.q, out.Val, tc.want)
			}
		})
	}
}
