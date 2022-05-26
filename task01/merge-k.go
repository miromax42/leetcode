package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// [[1,4,5],[1,3,4],[2,6]]

func main() {
	lists := []*ListNode{
		{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
	}

	fmt.Println(mergeKLists(lists))
}

func mergeKLists(lists []*ListNode) *ListNode {
	final := &ListNode{}

	for {

		var min *ListNode
		var position int

		for i := range lists {
			if lists[i] == nil {
				continue
			}
			if min == nil || min.Val > lists[i].Val {
				min = lists[i]
				position = i
				fmt.Println(min)
			}
		}

		if min == nil {
			break
		}

		final.Next = min
		lists[position] = lists[position].Next
	}

	return final.Next
}
