package structures_test

import (
	"testing"

	"github.com/olegshulyakov/leet-code.go/structures"
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
	in := []int{1}
	out := structures.NewLinkedList(in)

	if out.Val != 1 {
		t.Errorf("want out.Val = 1, got %d", out.Val)
	}
	if out.Next != nil {
		t.Error("want out.Next = nil")
	}
}

func TestNewLinkedListMultiple(t *testing.T) {
	in := []int{1, 2, 3}
	out := structures.NewLinkedList(in)

	if out.Val != 1 {
		t.Errorf("want out.Val = 1, got %d", out.Val)
	}
	if out.Next.Val != 2 {
		t.Errorf("want out.Next.Val = 2, got %d", out.Next.Val)
	}
	if out.Next.Next.Val != 3 {
		t.Errorf("want out.Next.Next.Val = 3, got %d", out.Next.Next.Val)
	}
	if out.Next.Next.Next != nil {
		t.Error("want out.Next.Next.Next = nil")
	}
}

func TestNewLinkedListEmpty(t *testing.T) {
	in := []int{}
	out := structures.NewLinkedList(in)

	if out != nil {
		t.Error("want out = nil")
	}
}

func TestToArraySingle(t *testing.T) {
	node := &structures.ListNode{Val: 1}
	out := node.ToArray()
	want := []int{1}
	if len(out) != len(want) {
		t.Fatalf("want length %d, got %d", len(want), len(out))
	}
	for i, v := range want {
		if out[i] != v {
			t.Errorf("want out[%d] = %d, got %d", i, v, out[i])
		}
	}
}

func TestToArrayMultiple(t *testing.T) {
	node := &structures.ListNode{Val: 2, Next: &structures.ListNode{Val: 1, Next: &structures.ListNode{Val: 3}}}
	out := node.ToArray()
	want := []int{2, 1, 3}
	if len(out) != len(want) {
		t.Fatalf("want length %d, got %d", len(want), len(out))
	}
	for i, v := range want {
		if out[i] != v {
			t.Errorf("want out[%d] = %d, got %d", i, v, out[i])
		}
	}
}

func TestToArrayNil(t *testing.T) {
	var node *structures.ListNode
	out := node.ToArray()
	if len(out) != 0 {
		t.Errorf("want empty array, got length %d", len(out))
	}
}

func TestToArrayEmpty(t *testing.T) {
	node := &structures.ListNode{}
	out := node.ToArray()
	want := []int{0}
	if len(out) != len(want) {
		t.Fatalf("want length %d, got %d", len(want), len(out))
	}
	for i, v := range want {
		if out[i] != v {
			t.Errorf("want out[%d] = %d, got %d", i, v, out[i])
		}
	}
}
