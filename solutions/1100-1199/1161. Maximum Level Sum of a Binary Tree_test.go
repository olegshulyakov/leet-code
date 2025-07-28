package main

import (
	"testing"

	"github.com/olegshulyakov/leet-code.go/structures"
)

/*

Given the root of a binary tree, the level of its root is 1, the level of its children is 2, and so on.

Return the smallest level x such that the sum of all the values of nodes at level x is maximal.


Example 1:

Input: root = [1,7,0,7,-8,null,null]
Output: 2
Explanation:
Level 1 sum = 1.
Level 2 sum = 7 + 0 = 7.
Level 3 sum = 7 + -8 = -1.
So we return the level with the maximum sum which is level 2.


Example 2:

Input: root = [989,null,10250,98693,-89388,null,null,null,-32127]
Output: 2


Constraints:

The number of nodes in the tree is in the range [1, 10^4].
-10^5 <= Node.val <= 10^5

*/

// maxLevelSum finds the level with the maximum sum of node values in a binary tree.
// Levels are numbered starting from 1 (root level).
//
// The algorithm uses breadth-first search (level-order traversal):
// 1. Start with the root level and track its sum
// 2. For each subsequent level, calculate the sum of all node values
// 3. Keep track of the level with the maximum sum
//
// Parameters:
//
//	root *TreeNode - Root of the binary tree
//
// Returns:
//
//	int - Level number (1-indexed) with the maximum sum
func maxLevelSum(root *structures.TreeNode) int {
	// Handle empty tree
	if root == nil {
		return 0
	}

	// Initialize variables
	maxSum := root.Val // Maximum sum found so far
	resultLevel := 1   // Level with maximum sum
	currentLevel := 1  // Current level being processed

	// Initialize queue with children of root for next level
	queue := []*structures.TreeNode{root.Left, root.Right}

	// Process each level
	for len(queue) > 0 {
		currentLevel++

		// Calculate sum of current level
		levelSum := 0
		for _, node := range queue {
			if node != nil {
				levelSum += node.Val
			}
		}

		// Update maximum if current level has higher sum
		if levelSum > maxSum {
			maxSum = levelSum
			resultLevel = currentLevel
		}

		// Prepare queue for next level
		newQueue := []*structures.TreeNode{}
		for _, node := range queue {
			if node == nil {
				continue
			}
			if node.Left != nil {
				newQueue = append(newQueue, node.Left)
			}
			if node.Right != nil {
				newQueue = append(newQueue, node.Right)
			}
		}
		queue = newQueue
	}

	return resultLevel
}

func TestMaxLevelSum(t *testing.T) {
	testCases := []struct {
		root []any
		want int
	}{
		{root: []any{1, 7, 0, 7, -8, nil, nil}, want: 2},
		{root: []any{989, nil, 10250, 98693, -89388, nil, nil, nil, -32127}, want: 2},
		{root: []any{}, want: 0},
		{root: []any{1}, want: 1},
	}
	for _, tc := range testCases {
		t.Run("1161. Maximum Level Sum of a Binary Tree", func(t *testing.T) {
			out := maxLevelSum(structures.NewTree(tc.root))
			if out != tc.want {
				t.Errorf("maxLevelSum(%v) = %v, want: %v", tc.root, out, tc.want)
			}
		})
	}
}
