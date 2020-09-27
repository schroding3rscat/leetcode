package main

import (
	"fmt"
)

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	lists := []*ListNode{
		{1, &ListNode{4, &ListNode{Val: 5}}},
		{1, &ListNode{3, &ListNode{Val: 4}}},
		{2, &ListNode{Val: 6}},
	}
	merged := mergeKLists(lists)
	printList(merged)
}

func printList(l *ListNode) {
	safe := 0
	for l != nil && safe < 100 {
		fmt.Printf("%v", l.Val)
		if l.Next != nil {
			fmt.Print(" -> ")
		}
		l = l.Next
		safe++
	}
	fmt.Print("\n")
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	merged := &ListNode{}
	head := merged
	idx := 0

	for {
		var min *ListNode
		for i := range lists {
			if lists[i] != nil && (min == nil || lists[i].Val < min.Val) {
				min = lists[i]
				idx = i
			}
		}

		if min == nil {
			return head.Next
		}

		merged.Next = &(*min)
		merged = merged.Next
		lists[idx] = lists[idx].Next
	}
}
