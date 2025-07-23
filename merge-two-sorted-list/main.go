package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 != nil && list2 == nil {
		return list1
	}

	if list1 == nil && list2 != nil {
		return list2
	}

	var newHead = &ListNode{}
	current := newHead

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else if list2 != nil {
		current.Next = list2
	}

	return newHead.Next
}

func print(head *ListNode) {
	outputs := []string{}
	current := head
	for current != nil {
		outputs = append(outputs, fmt.Sprint(current.Val))
		current = current.Next
	}
	fmt.Println(strings.Join(outputs, " -> "))
}

func main() {
	head1 := &ListNode{Val: 1}
	node11 := &ListNode{Val: 2}
	node12 := &ListNode{Val: 4}

	head2 := &ListNode{Val: 1}
	node21 := &ListNode{Val: 3}
	node22 := &ListNode{Val: 4}

	head1.Next = node11
	node11.Next = node12

	head2.Next = node21
	node21.Next = node22

	print(head1)
	print(head2)
	print(mergeTwoLists(head1, head2))
}
