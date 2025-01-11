package main

import (
	"fmt"
)

// Main Question in binary tree -
// What to do if there is a reparative number - Usually you skip
// In case number is reparative - ask interviewer , what to do , do you want to add to left or right or mark on the same number.

// BinaryTree represents a binary tree.
type BinaryTree struct {
	Root *Node
}

type Node struct {
	right *Node
	left  *Node
	val   int
}

// InsertWithIteration Iteration method
func (tree *BinaryTree) InsertWithIteration(value int) {
	tree.Root = insertNodeIterative(tree.Root, value)
}

func insertNodeIterative(head *Node, val int) *Node {
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

// InsertWithRecursion Recursive method
func (tree *BinaryTree) InsertWithRecursion(value int) {
	tree.Root = insertNodeRecursive(tree.Root, value)
}

func insertNodeRecursive(root *Node, value int) *Node {
	if root == nil {
		return &Node{val: value}
	}

	if value < root.val {
		root.left = insertNodeRecursive(root.left, value)
	} else {
		root.right = insertNodeRecursive(root.right, value)
	}

	return root
}

// visitPreOrder - RLR (Root - Left - Right)
func visitPreOrder(node *Node) []int {
	var vLeft, vRight []int
	res := []int{}
	res = append(res, node.val)

	if node.left != nil {
		vLeft = visitPreOrder(node.left)
	}
	res = append(res, vLeft...)

	if node.right != nil {
		vRight = visitPreOrder(node.right)
	}
	res = append(res, vRight...)
	return res
}

// Visit PostOrder - LRR (Left - Right - Root)
func visitPostOrder(node *Node) []int {
	var vLeft, vRight []int
	if node.left != nil {
		vLeft = visitPostOrder(node.left)
	}
	if node.right != nil {
		vRight = visitPostOrder(node.right)
	}
	res := []int{}
	res = append(res, vLeft...)
	res = append(res, vRight...)
	res = append(res, node.val)
	return res
}

// visitInOrder - LRR (Left - Root - Right)
func visitInOrder(node *Node) []int {
	var vLeft, vRight []int
	if node.left != nil {
		vLeft = visitInOrder(node.left)
	}
	res := []int{}
	res = append(res, vLeft...)
	res = append(res, node.val)
	if node.right != nil {
		vRight = visitInOrder(node.right)
	}
	res = append(res, vRight...)
	return res
}

// In - [2 9 10 11 23 54 65 67 89 90]
// Pre - [10 9 2 90 65 23 11 54 67 89]
// Post - [2 9 11 54 23 89 67 65 90 10]

func main() {
	arr := []int{10, 90, 65, 23, 11, 67, 89, 54, 9, 2, 65}

	treeR := &BinaryTree{}
	treeI := &BinaryTree{}

	for _, v := range arr {
		treeI.InsertWithIteration(v)
	}

	for _, v := range arr {
		treeR.InsertWithRecursion(v)
	}

	post1 := visitInOrder(treeI.Root)
	post2 := visitInOrder(treeR.Root)

	fmt.Println(post1)
	fmt.Println(post2)

	for i := range post1 {
		if post1[i] != post2[i] {
			fmt.Println(post1[i])
			fmt.Println(post2[i])

			fmt.Println("not same 1 : ", i)
		}
	}

	post1 = visitPreOrder(treeI.Root)
	post2 = visitPreOrder(treeR.Root)

	fmt.Println(post1)
	fmt.Println(post2)

	for i := range post1 {
		if post1[i] != post2[i] {
			fmt.Println("not same 2 : ", i)
		}
	}

	post1 = visitPostOrder(treeI.Root)
	post2 = visitPostOrder(treeR.Root)

	fmt.Println(post1)
	fmt.Println(post2)

	for i := range post1 {
		if post1[i] != post2[i] {
			fmt.Println("not same 3 : ", i)
		}
	}
}
