package main

import (
	"testing"

	"github.com/olegshulyakov/leet-code.go/structures"
)

/*

Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.

Basically, the deletion can be divided into two stages:

Search for a node to remove.
If the node is found, delete the node.


Example 1:

Input: root = [5,3,6,2,4,null,7], key = 3
Output: [5,4,6,2,null,null,7]
Explanation: Given key to delete is 3. So we find the node with value 3 and delete it.
One valid answer is [5,4,6,2,null,null,7], shown in the above BST.
Please notice that another valid answer is [5,2,6,null,4,null,7] and it's also accepted.


Example 2:

Input: root = [5,3,6,2,4,null,7], key = 0
Output: [5,3,6,2,4,null,7]
Explanation: The tree does not contain a node with value = 0.


Example 3:

Input: root = [], key = 0
Output: []


Constraints:

The number of nodes in the tree is in the range [0, 10^4].
-10^5 <= Node.val <= 10^5
Each node has a unique value.
root is a valid binary search tree.
-10^5 <= key <= 10^5


Follow up: Could you solve it with time complexity O(height of tree)?

*/

// deleteNode deletes a node with the specified key from the binary search tree.
// It maintains the BST property after deletion.
//
// The deletion process follows these cases:
//  1. Node not found: return the tree unchanged
//  2. Node has no children: simply remove the node
//  3. Node has one child: replace the node with its child
//  4. Node has two children: replace the node's value with its inorder successor
//     and delete the successor node
//
// Parameters:
//
//	root *TreeNode - Root of the binary search tree
//	key int - Value of the node to be deleted
//
// Returns:
//
//	*TreeNode - Root of the modified binary search tree
func deleteNode(root *structures.TreeNode, key int) *structures.TreeNode {
	// Base case: empty tree
	if root == nil {
		return root
	}

	// Key is smaller than root's value, delete from left subtree
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
		return root
	}

	// Key is greater than root's value, delete from right subtree
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
		return root
	}

	// Key is equal to root's value, this is the node to be deleted

	// Case 1: Node has no left child
	if root.Left == nil {
		return root.Right
	}

	// Case 2: Node has no right child
	if root.Right == nil {
		return root.Left
	}

	// Case 3: Node has both children
	// Find the inorder successor (smallest node in right subtree)
	minimal := root.Right
	for minimal.Left != nil {
		minimal = minimal.Left
	}

	// Replace the node's value with successor's value
	root.Val = minimal.Val

	// Delete the successor node
	root.Right = deleteNode(root.Right, minimal.Val)

	return root
}

func TestDeleteNode(t *testing.T) {
	testCases := []struct {
		nums []any
		key  int
		want []any
	}{
		{nums: []any{}, key: 0, want: []any{}},
		{nums: []any{5, 3, 6, 2, 4, nil, 7}, key: 0, want: []any{5, 3, 6, 2, 4, nil, 7}},
		{nums: []any{5, 3, 6, 2, 4, nil, 7}, key: 3, want: []any{5, 4, 6, 2, nil, nil, 7}},
		{nums: []any{0}, key: 0, want: []any{}},
		{nums: []any{50, 30, 70, nil, 40, 60, 80}, key: 50, want: []any{60, 30, 70, nil, 40, nil, 80}},
	}
	for _, tc := range testCases {
		t.Run("450. Delete Node in a BST", func(t *testing.T) {
			t.Logf("Delete Node in a BST: nums=%d, key=%d", tc.nums, tc.key)
			root := structures.NewTree(tc.nums)

			newRoot := deleteNode(root, tc.key)
			t.Logf("Delete Node in a BST: root:\n%v, newRoot:\n%v", root.Print(), newRoot.Print())
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
