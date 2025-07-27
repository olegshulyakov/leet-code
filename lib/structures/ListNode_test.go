package structures_test

import (
	"testing"

	"github.com/olegshulyakov/leet-code.go/lib/structures"
)

// TestListNodeCreation tests the creation of ListNode with various parameters.
func TestListNodeCreation(t *testing.T) {
	// Test creation with default values
	t.Run("should create a new ListNode with default values", func(t *testing.T) {
		node := &structures.ListNode{}
		if node.Val != 0 {
			t.Errorf("want node.Val = 0, got %d", node.Val)
		}
		if node.Next != nil {
			t.Error("want node.Next = nil")
		}
	})

	// Test creation with custom value and nil next
	t.Run("should create a new ListNode with custom value and null next", func(t *testing.T) {
		node := &structures.ListNode{Val: 5}
		if node.Val != 5 {
			t.Errorf("want node.Val = 5, got %d", node.Val)
		}
		if node.Next != nil {
			t.Error("want node.Next = nil")
		}
	})

	// Test creation with default value and custom next
	t.Run("should create a new ListNode with default value and custom next", func(t *testing.T) {
		node1 := &structures.ListNode{}
		node2 := &structures.ListNode{Val: 10}
		node1.Next = node2
		if node1.Val != 0 {
			t.Errorf("want node1.Val = 0, got %d", node1.Val)
		}
		if node2.Val != 10 {
			t.Errorf("want node2.Val = 10, got %d", node2.Val)
		}
		if node2.Next != nil {
			t.Error("want node2.Next = nil")
		}
	})
}

// TestNewLinkedList tests the NewLinkedList function with various inputs.
func TestNewLinkedList(t *testing.T) {
	// Test creation from single element
	t.Run("should create a linked list from a single element", func(t *testing.T) {
		nums := []int{1}
		result := structures.NewLinkedList(nums)
		if result.Val != 1 {
			t.Errorf("want result.Val = 1, got %d", result.Val)
		}
		if result.Next != nil {
			t.Error("want result.Next = nil")
		}
	})

	// Test creation from multiple elements
	t.Run("should create a linked list from multiple elements in reverse order", func(t *testing.T) {
		nums := []int{1, 2, 3}
		result := structures.NewLinkedList(nums)
		if result.Val != 1 {
			t.Errorf("want result.Val = 1, got %d", result.Val)
		}
		if result.Next.Val != 2 {
			t.Errorf("want result.Next.Val = 2, got %d", result.Next.Val)
		}
		if result.Next.Next.Val != 3 {
			t.Errorf("want result.Next.Next.Val = 3, got %d", result.Next.Next.Val)
		}
		if result.Next.Next.Next != nil {
			t.Error("want result.Next.Next.Next = nil")
		}
	})

	// Test creation from empty array
	t.Run("should return nil when creating a linked list from empty array", func(t *testing.T) {
		nums := []int{}
		result := structures.NewLinkedList(nums)
		if result != nil {
			t.Error("want result = nil")
		}
	})

	// Test creation from array with only one element (manual construction)
	t.Run("should create a linked list from array with only one element and no head", func(t *testing.T) {
		nums := []int{1}
		var head *structures.ListNode
		for i := len(nums) - 1; i >= 0; i-- {
			head = &structures.ListNode{Val: nums[i], Next: head}
		}
		if head.Val != 1 {
			t.Errorf("want head.Val = 1, got %d", head.Val)
		}
	})

	// Test creation from array with multiple elements
	t.Run("should create a linked list from array with multiple elements and non-empty head", func(t *testing.T) {
		nums := []int{2, 3, 4}
		head := structures.NewLinkedList(nums)
		if head.Val != 2 {
			t.Errorf("want head.Val = 2, got %d", head.Val)
		}
		if head.Next.Val != 3 {
			t.Errorf("want head.Next.Val = 3, got %d", head.Next.Val)
		}
		if head.Next.Next.Val != 4 {
			t.Errorf("want head.Next.Next.Val = 4, got %d", head.Next.Next.Val)
		}
	})
}

// TestToArray tests the ToArray method with various inputs.
func TestToArray(t *testing.T) {
	// Test conversion of single-element linked list
	t.Run("should create an array from a single-element linked list", func(t *testing.T) {
		node := &structures.ListNode{Val: 1}
		result := node.ToArray()
		want := []int{1}
		if len(result) != len(want) {
			t.Fatalf("want length %d, got %d", len(want), len(result))
		}
		for i, v := range want {
			if result[i] != v {
				t.Errorf("want result[%d] = %d, got %d", i, v, result[i])
			}
		}
	})

	// Test conversion of multi-element linked list
	t.Run("should create an array from a multi-element linked list in reverse order", func(t *testing.T) {
		node := &structures.ListNode{Val: 2, Next: &structures.ListNode{Val: 1, Next: &structures.ListNode{Val: 3}}}
		result := node.ToArray()
		want := []int{2, 1, 3}
		if len(result) != len(want) {
			t.Fatalf("want length %d, got %d", len(want), len(result))
		}
		for i, v := range want {
			if result[i] != v {
				t.Errorf("want result[%d] = %d, got %d", i, v, result[i])
			}
		}
	})

	// Test conversion of nil linked list
	t.Run("should return empty array when passing nil to ToArray function", func(t *testing.T) {
		var node *structures.ListNode
		result := node.ToArray()
		if len(result) != 0 {
			t.Errorf("want empty array, got length %d", len(result))
		}
	})

	// Test conversion of empty linked list (node with default value)
	t.Run("should create an array from empty linked list", func(t *testing.T) {
		node := &structures.ListNode{}
		result := node.ToArray()
		want := []int{0}
		if len(result) != len(want) {
			t.Fatalf("want length %d, got %d", len(want), len(result))
		}
		for i, v := range want {
			if result[i] != v {
				t.Errorf("want result[%d] = %d, got %d", i, v, result[i])
			}
		}
	})
}
