package main

import (
	"testing"

	"github.com/olegshulyakov/leet-code.go/structures"
)

/*

You are given the root of a binary tree.

A ZigZag path for a binary tree is defined as follow:

Choose any node in the binary tree and a direction (right or left).
If the current direction is right, move to the right child of the current node; otherwise, move to the left child.
Change the direction from right to left or from left to right.
Repeat the second and third steps until you can't move in the tree.
Zigzag length is defined as the number of nodes visited - 1. (A single node has a length of 0).

Return the longest ZigZag path contained in that tree.


Example 1:

Input: root = [1,null,1,1,1,null,null,1,1,null,1,null,null,null,1]
Output: 3
Explanation: Longest ZigZag path in blue nodes (right -> left -> right).


Example 2:

Input: root = [1,1,1,null,1,null,null,1,1,null,1]
Output: 4
Explanation: Longest ZigZag path in blue nodes (left -> right -> left -> right).


Example 3:

Input: root = [1]
Output: 0


Constraints:

The number of nodes in the tree is in the range [1, 5 * 10^4].
1 <= Node.val <= 100

*/

// longestZigZag finds the length of the longest zigzag path in a binary tree.
// A zigzag path alternates between left and right moves.
// The length is defined as the number of nodes visited minus 1.
//
// The algorithm:
// 1. For each node, try starting a zigzag path in both directions (left and right)
// 2. For each starting direction, follow the zigzag pattern as far as possible
// 3. Track the maximum length found across all nodes and starting directions
//
// Parameters:
//
//	root *TreeNode - Root of the binary tree
//
// Returns:
//
//	int - Length of the longest zigzag path
func longestZigZag(root *structures.TreeNode) int {
	// Handle empty tree
	if root == nil {
		return 0
	}

	maxLength := 0
	queue := []*structures.TreeNode{root}

	// Process each node as a potential starting point
	for len(queue) > 0 {
		// Get current node (pop from end of slice)
		cursor := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		// Try both starting directions: left (true) and right (false)
		for _, startLeft := range []bool{true, false} {
			count := 0
			node := cursor
			goLeft := startLeft

			// Follow zigzag path as far as possible
			for node != nil {
				if goLeft {
					node = node.Left
				} else {
					node = node.Right
				}
				goLeft = !goLeft // Alternate direction
				count++
			}

			// Decrement count because length is nodes - 1
			count--

			// Update maximum length if current path is longer
			if count > maxLength {
				maxLength = count
			}
		}

		// Add children to queue for processing as starting points
		if cursor.Left != nil {
			queue = append(queue, cursor.Left)
		}
		if cursor.Right != nil {
			queue = append(queue, cursor.Right)
		}
	}

	return maxLength
}

func TestLongestZigZag(t *testing.T) {
	testCases := []struct {
		in   []any
		want int
	}{
		{in: []any{1, nil, 1, 1, 1, nil, nil, 1, 1, nil, 1, nil, nil, nil, 1}, want: 3},
		{in: []any{1, 1, 1, nil, 1, nil, nil, 1, 1, nil, 1}, want: 4},
		{in: []any{1}, want: 0},
	}
	for _, tc := range testCases {
		t.Run("1372. Longest ZigZag Path in a Binary Tree", func(t *testing.T) {
			out := longestZigZag(structures.NewTree(tc.in))
			if out != tc.want {
				t.Errorf("longestZigZag(%v) = %v, want: %v", tc.in, out, tc.want)
			}
		})
	}
}
