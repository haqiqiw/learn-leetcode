package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// main idea
// nodes = nil/dummy, 1, 2, 3, 4, 5
// 1 / 3 = +1 for slow / +n for fast
// 2 / 4 = cotinue +1 for both
// 3 / 5

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	dummyHead := &ListNode{Next: head}
	slow, fast := dummyHead, dummyHead
	initial := true

	for fast != nil {
		if initial {
			for i := 0; i <= n; i++ {
				fast = fast.Next
			}
			initial = false
		} else {
			fast = fast.Next
		}

		if fast == nil {
			slow.Next = slow.Next.Next
		} else {
			slow = slow.Next
		}
	}

	return dummyHead.Next
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

	print(removeNthFromEnd(head, 2))
	print(removeNthFromEnd(&ListNode{Val: 1}, 1))
	print(removeNthFromEnd(&ListNode{Val: 1, Next: &ListNode{Val: 2}}, 1))
}
