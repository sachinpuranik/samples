package main

import (
	"fmt"
)

type Node struct {
	left  *Node
	right *Node
	val   int
}

func InsertNode(head *Node, newNode *Node) *Node {
	if head == nil {
		head = newNode
		return head
	}

	if head.val > newNode.val && head.left != nil {
		InsertNode(head.left, newNode)
	}

	if head.val < newNode.val && head.right != nil {
		InsertNode(head.right, newNode)
	}

	if head.val > newNode.val && head.left == nil {
		head.left = newNode
		return head
	}

	if head.val < newNode.val && head.right == nil {
		head.right = newNode
		return head
	}

	return head
}

var head *Node

func VisitInOrder(head *Node) {
	if head.left != nil {
		VisitInOrder(head.left)
	}
	fmt.Println(head.val)
	if head.right != nil {
		VisitInOrder(head.right)
	}
}

func main() {
	var arr []int = []int{6, 10, 58, 32, 48, 7, 98}
	for _, val := range arr {
		newNode := &Node{val: val}
		head = InsertNode(head, newNode)
	}

	VisitInOrder(head)
}
