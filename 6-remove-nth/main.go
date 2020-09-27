package main

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, nil}}}}}}}
	_ = removeNthFromEnd(&list, 3)

	list2 := ListNode{1, &ListNode{2, nil}}
	_ = removeNthFromEnd(&list2, 2)
}

// 1 - 2 - 3 - 4 - !5! - 6 - 7
// n = 3
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		if n == 1 {
			return nil
		}
		return head
	}

	h := &ListNode{Next: head}
	first := h
	last := h

	for i := 1; first != nil && i <= n+1; i++ {
		first = first.Next
	}

	for first != nil {
		last = last.Next
		first = first.Next
	}

	last.Next = last.Next.Next

	return h.Next
}
