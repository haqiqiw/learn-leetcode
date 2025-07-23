package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow, fast, middle := head, head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		middle = slow
	}

	return middle
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
	node5 := &ListNode{Val: 6}

	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	print(middleNode(head))
}
