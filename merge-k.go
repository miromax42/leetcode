package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// [[1,4,5],[1,3,4],[2,6]]

func main() {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}

	fmt.Println(addTwoNumbers(list1, list2))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := &ListNode{}
	current := sum
	add := 0

	for {
		var cur1, cur2 int

		if l1 == nil && l2 == nil {
			return sum.Next
		} else {
			current.Next = &ListNode{}
			current = current.Next
		}

		if l1 != nil {
			fmt.Println(l1.Val)
			cur1 = l1.Val
			l1 = l1.Next

		}

		if l2 != nil {
			cur2 = l2.Val
			l2 = l2.Next
		}

		fullSum := cur1 + cur2 + add
		digit := fullSum / 10
		add = fullSum % 10

		current.Val = digit

	}
}
