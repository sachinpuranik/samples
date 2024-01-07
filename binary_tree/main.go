package main

import (
	"fmt"
)

type Node struct {
	right *Node
	left  *Node
	val   int
}

func InsertNode(head *Node, val int) *Node {
	node := &Node{val: val}
	mover := head
	if head == nil {
		head = node
		return head
	}
	for {
		if mover.val >= node.val && mover.left != nil {
			mover = mover.left
			continue
		}
		if mover.val >= node.val {
			mover.left = node
			break
		}
		if mover.val < node.val && mover.right != nil {
			mover = mover.right
			continue
		}
		if mover.val < node.val {
			mover.right = node
			break
		}
	}
	return head
}

func CreateTree(arr []int) *Node {
	var head *Node
	for _, v := range arr {
		head = InsertNode(head, v)
	}
	return head
}

func VisitTree(node *Node) {
	if node.left != nil {
		VisitTree(node.left)
	}
	if node.right != nil {
		VisitTree(node.right)
	}
	fmt.Println(node.val)
}

func main() {
	arr := []int{10, 90, 65, 23, 11, 67, 89, 54, 65}
	tree := CreateTree(arr)
	VisitTree(tree)
}
