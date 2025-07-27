// Package structures provides data structures.
package structures

// ListNode represents a node in a singly linked list.
// Each node contains an integer value and a pointer to the next node in the list.
type ListNode struct {
	Val  int       // Val holds the integer value stored in this node
	Next *ListNode // Next points to the subsequent node in the list (nil if this is the last node)
}

// NewLinkedList creates a singly linked list from an array of integers.
// The function constructs the list by iterating through the array from right to left,
// building the list in reverse order to maintain the correct sequence.
//
// Example:
//
//	nums := []int{1, 2, 3, 4}
//	list := NewLinkedList(nums) // Creates list: 1 -> 2 -> 3 -> 4 -> nil
//
// Parameters:
//
//	nums []int - Array of integers to convert into a linked list
//
// Returns:
//
//	*ListNode - Pointer to the head node of the created linked list (nil if input array is empty)
func NewLinkedList(nums []int) *ListNode {
	// Return nil for empty input array
	if len(nums) == 0 {
		return nil
	}

	// Declare head
	var head *ListNode

	// Iterate through the array from right to left (reverse order)
	// This allows us to build the linked list in the correct order
	for i := len(nums) - 1; i >= 0; i-- {
		// Create new node with current value and link it to the existing list
		// The new node becomes the new head of the list
		head = &ListNode{Val: nums[i], Next: head}
	}

	// Return pointer to the head of the constructed linked list
	return head
}

// ToArray converts the linked list to an array of integers.
// The function traverses the linked list from head to tail, collecting
// each node's value into a slice.
//
// Example:
//
//	list := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
//	arr := list.ToArray() // Returns []int{1, 2, 3}
//
// Returns:
//
//	[]int - Slice containing all values from the linked list in order
func (list *ListNode) ToArray() []int {
	// Return empty slice for nil list
	if list == nil {
		return []int{}
	}

	// Initialize result slice to store list values
	res := []int{}

	// Initialize pointer to traverse the list starting from head
	pointer := list

	// Traverse the list until we reach the end (nil)
	for pointer != nil {
		// Append current node's value to result slice
		res = append(res, pointer.Val)
		// Move to next node in the list
		pointer = pointer.Next
	}

	// Return the array representation of the linked list
	return res
}
