package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var buffer, prev *ListNode
	current := head

	for current != nil {
		buffer = current.Next
		current.Next = prev

		prev = current
		current = buffer

	}

	return prev
}
