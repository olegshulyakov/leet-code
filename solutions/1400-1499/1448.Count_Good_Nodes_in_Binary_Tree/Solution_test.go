package main

import (
	"math"
	"testing"

	"github.com/olegshulyakov/leet-code.go/structures"
)

/*

Given a binary tree root, a node X in the tree is named good if in the path from root to X there are no nodes with a value greater than X.

Return the number of good nodes in the binary tree.


Example 1:

Input: root = [3,1,4,3,null,1,5]
Output: 4
Explanation: Nodes in blue are good.
Root Node (3) is always a good node.
Node 4 -> (3,4) is the maximum value in the path starting from the root.
Node 5 -> (3,4,5) is the maximum value in the path
Node 3 -> (3,1,3) is the maximum value in the path.


Example 2:

Input: root = [3,3,null,4,2]
Output: 3
Explanation: Node 2 -> (3, 3, 2) is not good, because "3" is higher than it.


Example 3:

Input: root = [1]
Output: 1
Explanation: Root is considered as good.


Constraints:

The number of nodes in the binary tree is in the range [1, 10^5].
Each node's value is between [-10^4, 10^4].

*/

// goodNodes counts the number of "good" nodes in a binary tree.
// A node X is considered good if in the path from root to X,
// there are no nodes with a value greater than X.
//
// The algorithm uses a recursive approach:
// 1. For each node, track the maximum value seen so far in the path
// 2. If current node's value >= max value so far, it's a good node
// 3. Recursively check left and right subtrees with updated max value
//
// Parameters:
//
//	root *TreeNode - Root of the binary tree
//
// Returns:
//
//	int - Number of good nodes in the tree
func goodNodes(root *structures.TreeNode) int {
	// Handle empty tree
	if root == nil {
		return 0
	}

	// Start recursive check with negative infinity as initial max value
	return checkSubTree(root, math.MinInt)
}

// checkSubTree recursively counts good nodes in a subtree.
// It tracks the maximum value encountered in the path from root to current node.
//
// Parameters:
//
//	root *TreeNode - Current node being checked
//	maxValue int - Maximum value in the path from original root to current node
//
// Returns:
//
//	int - Number of good nodes in the subtree rooted at 'root'
func checkSubTree(root *structures.TreeNode, maxValue int) int {
	// Update max value for this path
	maxVal := root.Val
	if maxVal <= maxValue {
		maxVal = maxValue
	}

	// Count current node as good if its value >= max value in path
	count := 0
	if root.Val >= maxValue {
		count = 1
	}

	// Recursively count good nodes in left and right subtrees
	leftCount := 0
	if root.Left != nil {
		leftCount = checkSubTree(root.Left, maxVal)
	}
	rightCount := 0
	if root.Right != nil {
		rightCount = checkSubTree(root.Right, maxVal)
	}

	return count + leftCount + rightCount
}

func TestGoodNodes(t *testing.T) {
	testCases := []struct {
		root []any
		want int
	}{
		{root: []any{3, 1, 4, 3, nil, 1, 5}, want: 4},
		{root: []any{3, 3, nil, 4, 2}, want: 3},
		{root: []any{1}, want: 1},
		{root: []any{9, nil, 3, 6}, want: 1},
	}
	for _, tc := range testCases {
		t.Run("1448. Count Good Nodes in Binary Tree", func(t *testing.T) {
			out := goodNodes(structures.NewTree(tc.root))

			if out != tc.want {
				t.Errorf("goodNodes(%v) = %v, want: %v", tc.root, out, tc.want)
			}
		})
	}
}
