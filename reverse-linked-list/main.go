package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var current, next, prev *ListNode
	current = head
	prev = nil

	for current != nil {
		next = current.Next

		current.Next = prev
		prev = current

		current = next
	}
	current = prev

	return current
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
	node1 := &ListNode{Val: 2}
	node2 := &ListNode{Val: 3}
	node3 := &ListNode{Val: 4}
	node4 := &ListNode{Val: 5}

	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	print(head)
	print(reverseList(head))
}
