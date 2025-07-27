package structures_test

import (
	"testing"

	"github.com/olegshulyakov/leet-code.go/lib/structures"
)

func TestListNodeDefaultValues(t *testing.T) {
	node := &structures.ListNode{}

	if node.Val != 0 {
		t.Errorf("want node.Val = 0, got %d", node.Val)
	}
	if node.Next != nil {
		t.Error("want node.Next = nil")
	}
}

func TestListNodeValueAndNextNil(t *testing.T) {
	node := &structures.ListNode{Val: 5}

	if node.Val != 5 {
		t.Errorf("want node.Val = 5, got %d", node.Val)
	}
	if node.Next != nil {
		t.Error("want node.Next = nil")
	}
}

func TestListNodeValueAndNextNotNil(t *testing.T) {
	node1 := &structures.ListNode{}
	node2 := &structures.ListNode{Val: 10}
	node1.Next = node2

	if node1.Val != 0 {
		t.Errorf("want node1.Val = 0, got %d", node1.Val)
	}
	if node1.Next == nil {
		t.Error("want node1.Next != nil")
	}
	if node1.Next.Val != 10 {
		t.Errorf("want node2.Val = 10, got %d", node2.Val)
	}
	if node2.Next != nil {
		t.Error("want node2.Next = nil")
	}
}

func TestNewLinkedListSingle(t *testing.T) {
	nums := []int{1}
	result := structures.NewLinkedList(nums)

	if result.Val != 1 {
		t.Errorf("want result.Val = 1, got %d", result.Val)
	}
	if result.Next != nil {
		t.Error("want result.Next = nil")
	}
}

func TestNewLinkedListMultiple(t *testing.T) {
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
}

func TestNewLinkedListEmpty(t *testing.T) {
	nums := []int{}
	result := structures.NewLinkedList(nums)

	if result != nil {
		t.Error("want result = nil")
	}
}

func TestToArraySingle(t *testing.T) {
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
}

func TestToArrayMultiple(t *testing.T) {
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
}

func TestToArrayNil(t *testing.T) {
	var node *structures.ListNode
	result := node.ToArray()
	if len(result) != 0 {
		t.Errorf("want empty array, got length %d", len(result))
	}
}

func TestToArrayEmpty(t *testing.T) {
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
}
