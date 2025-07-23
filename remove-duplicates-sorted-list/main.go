package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	current := head
	for current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
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
	head := &ListNode{Val: 1}
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}

	head.Next = node1
	node1.Next = node2

	print(head)
	deleteDuplicates(head)
	print(head)
}
