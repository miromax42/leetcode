package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// [[1,4,5],[1,3,4],[2,6]]

func main() {
	lists := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}

	fmt.Println(hasCycle(lists))
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	tortoise := head
	hair := head.Next

	for {
		if hair == nil {
			return false
		}

		if tortoise == hair {
			return true
		}

		tortoise = tortoise.Next
		hair = hair.Next

		if hair == nil {
			return false
		}

		hair = hair.Next

	}
}
