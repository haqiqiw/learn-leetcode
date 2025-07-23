package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	var slow, fast *ListNode

	slow = head
	fast = head

	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
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
	head := &ListNode{Val: 3}
	node1 := &ListNode{Val: 2}
	node2 := &ListNode{Val: 0}
	node3 := &ListNode{Val: 4}

	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node2

	fmt.Println(hasCycle(head))

	head2 := &ListNode{Val: 1}
	fmt.Println(hasCycle(head2))
}
