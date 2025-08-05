package main

import (
	"testing"

	"github.com/olegshulyakov/leet-code.go/structures"
)

/*

Given the root of a binary tree and an integer targetSum, return the number of paths where the sum of the values along the path equals targetSum.

The path does not need to start or end at the root or a leaf, but it must go downwards (i.e., traveling only from parent nodes to child nodes).


Example 1:

Input: root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
Output: 3
Explanation: The paths that sum to 8 are shown.


Example 2:

Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
Output: 3


Constraints:

The number of nodes in the tree is in the range [0, 1000].
-10^9 <= Node.val <= 10^9
-1000 <= targetSum <= 1000

*/

// pathSum counts the number of paths in a binary tree where the sum of node values equals targetSum.
// A path can start and end at any node, but must go downwards (parent to child only).
//
// The algorithm uses a double recursive approach:
// 1. For each node, count paths starting from that node (subSum function)
// 2. Recursively check all nodes as potential starting points
//
// Parameters:
//
//	root *TreeNode - Root of the binary tree
//	targetSum int - Target sum to find in paths
//
// Returns:
//
//	int - Number of paths with sum equal to targetSum
func pathSum(root *structures.TreeNode, targetSum int) int {
	// Base case: empty tree has 0 paths
	if root == nil {
		return 0
	}

	// Count paths starting from current node plus paths in subtrees
	return subSum(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

// subSum counts the number of paths starting from the given node where the sum equals target.
// This is a helper function that explores paths downward from a specific starting node.
//
// Parameters:
//
//	node *TreeNode - Starting node for path exploration
//	target int - Remaining target sum to reach
//
// Returns:
//
//	int - Number of valid paths starting from the given node
func subSum(node *structures.TreeNode, target int) int {
	// Base case: nil node contributes 0 paths
	if node == nil {
		return 0
	}

	// Check if current node completes a valid path
	count := 0
	if node.Val == target {
		count = 1
	}

	// Update target and explore children
	newTarget := target - node.Val
	leftCount := 0
	rightCount := 0

	if node.Left != nil {
		leftCount = subSum(node.Left, newTarget)
	}
	if node.Right != nil {
		rightCount = subSum(node.Right, newTarget)
	}

	return count + leftCount + rightCount
}

func TestPathSum(t *testing.T) {
	name := "437. Path Sum III"
	testCases := []struct {
		node   []any
		target int
		want   int
	}{
		{node: []any{10, 5, -3, 3, 2, nil, 11, 3, -2, nil, 1}, target: 8, want: 3},
		{node: []any{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1}, target: 22, want: 3},
		{node: []any{}, target: 1, want: 0},
		{node: []any{-2, nil, -3}, target: -5, want: 1},
		{node: []any{1, -2, -3, 1, 3, -2, nil, -1}, target: -1, want: 4},
	}
	for _, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			out := pathSum(structures.NewTree(tc.node), tc.target)
			if out != tc.want {
				t.Errorf("pathSum(%v, %v) = %v, want: %v", tc.node, tc.target, out, tc.want)
			}
		})
	}
}
