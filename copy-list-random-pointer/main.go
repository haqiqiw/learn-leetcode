package main

import (
	"fmt"
)

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	headMap := map[*Node]*Node{}

	current := head
	for current != nil {
		headMap[current] = &Node{Val: current.Val}

		current = current.Next
	}

	current = head
	for current != nil {
		newNode := headMap[current]

		newNode.Next = headMap[current.Next]
		newNode.Random = headMap[current.Random]

		current = current.Next
	}

	return headMap[head]
}

func print(head *Node) {
	outputs := [][]string{}
	current := head
	for current != nil {
		r := "null"
		if current.Random != nil {
			r = fmt.Sprint(current.Random.Val)
		}
		outputs = append(outputs, []string{fmt.Sprint(current.Val), r})
		current = current.Next
	}
	fmt.Println(outputs)
}

func main() {
	head := &Node{Val: 7}
	node1 := &Node{Val: 13}
	node2 := &Node{Val: 11}
	node3 := &Node{Val: 10}
	node4 := &Node{Val: 1}

	head.Next = node1
	head.Random = nil
	node1.Next = node2
	node1.Random = head
	node2.Next = node3
	node2.Random = node4
	node3.Next = node4
	node3.Random = node2
	node4.Next = nil
	node4.Random = head

	print(head)
	print(copyRandomList(head))
}
