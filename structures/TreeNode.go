// Package structures provides data structures.
package structures

import (
	"fmt"
	"strings"
)

// TreeNode defines a node in the binary tree.
// Each node contains an integer value and pointers to left and right children.
type TreeNode struct {
	Val   int       // Val holds the integer value of the node
	Left  *TreeNode // Left points to the left child node (nil if no child)
	Right *TreeNode // Right points to the right child node (nil if no child)
}

// NewTree generates a binary tree from an array representation.
// The array follows level-order traversal where 0 represents a null node.
// This function uses breadth-first construction to build the tree.
//
// Example:
//
//	arr := []int{1, 2, 3, 0, 4, 5, 0}
//	tree := NewTree(arr) // Creates tree with root=1, left=2, right=3, etc.
//
// Parameters:
//
//	arr []int - Array representation of the binary tree in level-order
//
// Returns:
//
//	*TreeNode - Root node of the constructed binary tree (nil if array is empty)
func NewTree(arr []any) *TreeNode {
	// Return nil for empty array
	if len(arr) == 0 {
		return nil
	}

	// Create root node with first element
	val, _ := arr[0].(int)
	root := &TreeNode{Val: val}

	// Initialize queue with root node for breadth-first construction
	queue := []*TreeNode{root}

	// Start processing from second element (index 1)
	index := 1

	// Continue until queue is empty or all array elements are processed
	for len(queue) > 0 && index < len(arr) {
		// Dequeue the front node
		node := queue[0]
		queue = queue[1:]

		// Process left child
		if index < len(arr) && arr[index] != nil {
			val, _ = arr[index].(int)
			// Create new node and link it as left child
			node.Left = &TreeNode{Val: val}
			// Enqueue the new node for processing its children
			queue = append(queue, node.Left)
		}
		// Move to next array element
		index++

		// Process right child
		if index < len(arr) && arr[index] != nil {
			val, _ = arr[index].(int)
			// Create new node and link it as right child
			node.Right = &TreeNode{Val: val}
			// Enqueue the new node for processing its children
			queue = append(queue, node.Right)
		}
		// Move to next array element
		index++
	}

	// Return the root of the constructed tree
	return root
}

// Print returns a string representation of the binary tree in a structured,
// visual format. The tree is displayed with branches showing the hierarchy.
//
// Example output:
// ├── 1
// │   ├── 2
// │   │   └── 4
// │   └── 3
// │       ├── 5
//
// Returns:
//
//	string - Visual representation of the tree structure
func (node *TreeNode) Print() string {
	// Return empty string for nil node
	if node == nil {
		return ""
	}

	// Call helper function starting at level 0, not last child
	return printHelper(node, 0, false)
}

// printHelper is a recursive helper function that builds the visual tree representation.
//
// Parameters:
//
//	node *TreeNode - Current node being processed
//	level int - Current depth level in the tree (used for indentation)
//	isLast bool - Whether this node is the last child of its parent
//
// Returns:
//
//	string - Formatted string representation of subtree rooted at node
func printHelper(node *TreeNode, level int, isLast bool) string {
	// Return empty string for nil node
	if node == nil {
		return ""
	}

	// Determine prefix based on whether this is the last child
	var prefix string
	if isLast {
		prefix = "└── " // Last child uses different branch symbol
	} else {
		prefix = "├── " // Non-last child uses standard branch symbol
	}

	// Build current node's line with proper indentation
	const indentSize = 4
	output := fmt.Sprintf("%s%s%d\n",
		strings.Repeat(" ", level*indentSize), // Indentation based on level
		prefix,                                // Branch symbol
		node.Val)                              // Node value

	// Recursively process children if they exist
	if node.Left != nil || node.Right != nil {
		// Process left child (not last)
		output += printHelper(node.Left, level+1, false)
		// Process right child (last)
		output += printHelper(node.Right, level+1, true)
	}

	return output
}

// ToArray converts the binary tree back to its array representation.
// The conversion follows level-order traversal where 0 represents null nodes.
// Trailing zeros (representing null nodes at the end) are removed.
//
// Returns:
//
//	[]int - Array representation of the binary tree in level-order traversal
func (node *TreeNode) ToArray() []any {
	// Return empty slice for nil tree
	if node == nil {
		return []any{}
	}

	// Initialize result slice and queue for level-order traversal
	res := []any{}
	queue := []*TreeNode{node}

	// Process nodes in level-order using BFS
	for len(queue) > 0 {
		// Dequeue front node
		current := queue[0]
		queue = queue[1:]

		if current != nil {
			// Add node value to result
			res = append(res, current.Val)
			// Enqueue children for processing (nil children are also enqueued)
			queue = append(queue, current.Left, current.Right)
		} else {
			// Add nil nodes
			res = append(res, nil)
		}
	}

	// Remove trailing nil values
	for len(res) > 0 {
		if res[len(res)-1] == nil {
			res = res[:len(res)-1]
		} else {
			break
		}
	}

	// Return the array representation
	return res
}
