package main

import (
	"testing"

	"github.com/olegshulyakov/leet-code.go/structures"
)

/*

Consider all the leaves of a binary tree, from left to right order, the values of those leaves form a leaf value sequence.


For example, in the given tree above, the leaf value sequence is (6, 7, 4, 9, 8).

Two binary trees are considered leaf-similar if their leaf value sequence is the same.

Return true if and only if the two given trees with head nodes root1 and root2 are leaf-similar.


Example 1:

Input: root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
Output: true


Example 2:

Input: root1 = [1,2,3], root2 = [1,3,2]
Output: false


Constraints:

The number of nodes in each tree will be in the range [1, 200].
Both of the given trees will have values in the range [0, 200].

*/

// leafSimilar determines if two binary trees are leaf-similar.
// Two trees are leaf-similar if their leaf value sequences (from left to right) are identical.
//
// The algorithm:
// 1. Extract leaf values from both trees in left-to-right order
// 2. Compare the two sequences for equality
//
// Parameters:
//
//	root1 *structures.TreeNode - Root of the first binary tree
//	root2 *structures.TreeNode - Root of the second binary tree
//
// Returns:
//
//	bool - True if trees are leaf-similar, false otherwise
func leafSimilar(root1 *structures.TreeNode, root2 *structures.TreeNode) bool {
	// Extract leaf sequences from both trees
	leafs1 := toLeafArray(root1)
	leafs2 := toLeafArray(root2)

	if len(leafs1) != len(leafs2) {
		return false
	}
	for i := range leafs1 {
		if leafs1[i] != leafs2[i] {
			return false
		}
	}
	return true
}

// toLeafArray extracts the leaf values from a binary tree in left-to-right order.
// A leaf is a node with no children (both left and right are nil).
//
// Parameters:
//
//	node *structures.TreeNode - Root of the tree to extract leaves from
//
// Returns:
//
//	[]int - Slice containing leaf values in left-to-right order
func toLeafArray(node *structures.TreeNode) []int {
	// Handle nil tree
	if node == nil {
		return []int{}
	}

	// Use stack for depth-first traversal (pre-order)
	queue := []*structures.TreeNode{node}
	res := []int{}

	// Process nodes
	for len(queue) > 0 {
		// Pop node from front of queue
		currentNode := queue[0]
		queue = queue[1:]

		// Skip nil nodes
		if currentNode == nil {
			continue
		}

		// Check if current node is a leaf
		if currentNode.Left == nil && currentNode.Right == nil {
			res = append(res, currentNode.Val)
		} else {
			// Add children to queue (left first for correct order)
			// Add to front of queue to maintain proper traversal order
			newQueue := []*structures.TreeNode{currentNode.Left, currentNode.Right}
			queue = append(newQueue, queue...)
		}
	}

	// Remove trailing zero values
	for len(res) > 0 && res[len(res)-1] == 0 {
		res = res[:len(res)-1]
	}

	return res
}

func TestLeafSimilar(t *testing.T) {
	testCases := []struct {
		root1 []any
		root2 []any
		want  bool
	}{
		{root1: []any{}, root2: []any{}, want: true},
		{root1: []any{3, 5, 1, 6, 2, 9, 8, nil, nil, 7, 4}, root2: []any{3, 5, 1, 6, 7, 4, 2, nil, nil, nil, nil, nil, nil, 9, 8}, want: true},
		{root1: []any{1, 2, 3}, root2: []any{1, 3, 2}, want: false},
		{root1: []any{1, 2}, root2: []any{2, 2}, want: true},
		{root1: []any{1}, root2: []any{2}, want: false},
		{root1: []any{1}, root2: []any{1}, want: true},
	}
	for _, tc := range testCases {
		t.Run("872. Leaf-Similar Trees", func(t *testing.T) {
			out := leafSimilar(structures.NewTree(tc.root1), structures.NewTree(tc.root2))
			if out != tc.want {
				t.Errorf("leafSimilar(%v, %v) = %v, want: %v", tc.root1, tc.root2, out, tc.want)
			}
		})
	}
}
