package leetcode

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	final := &ListNode{}

	min := lists[0]
	for i := range lists {
		if lists[i] == nil {
			continue
		}
		if min.Val > lists[i].Val {
			min = lists[i]
		}
	}

	final.Next = min
	min = min.Next

	return final.Next
}
