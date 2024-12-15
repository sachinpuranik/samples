package main

import (
	"fmt"
)

// Node -
type Node struct {
	Next  *Node
	Value int
}

// basic algo as below.
// 1. first step if head or head.next is nil , list do not have loop.
// 2. two pointers slow and fast walk behind each other.slow walks 1 step vs fast walk 2 steps always. so initilize two pointers to head and head.next
// 3. Fast will always approach the last node fast , else it will reach slow ptr position, if loop exist.
// 4. loop when fast is not null and fast next is also not nil.
//  5. check if slow and fast are same, then break as loop exist.
//  6. since we have already checked that fast and fast.next are not nil while starting loop , we can safely do fast.next and slow.next.
//
// Continue through 4-6 till loop is found or pointer reaches nil. and you know the result.
func CheckCycle(head *Node) bool {
	// Handle empty or single-node lists directly:
	if head == nil || head.Next == nil {
		return false
	}

	// Use descriptive variable names:
	slowPointer := head
	fastPointer := head.Next

	// Loop until fastPointer reaches the end or a cycle is detected:
	for fastPointer != nil && fastPointer.Next != nil {
		if fastPointer == slowPointer {
			return true // Cycle detected
		}

		// Move pointers at different speeds:
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next
	}

	return false // No cycle found
}

// InsertNode -
func InsertNode(node, head *Node) *Node {
	node.Next = head
	head = node
	return head
}

// CreateCycle -
func CreateCycle(head *Node) *Node {
	if head == nil {
		return nil
	}
	last := head
	for last.Next != nil {
		last = last.Next
	}
	last.Next = head
	return head
}

func main() {
	var head *Node

	// Test 1 head = nil
	head = nil
	if CheckCycle(head) == true {
		fmt.Println("has Cycle")
	} else {
		fmt.Println("No Cycle")
	}

	// Test 2 Linked list with 1 node
	head = nil
	for i := 0; i < 1; i++ {
		node := &Node{Value: i + 1}
		head = InsertNode(node, head)
	}

	// Test 2.1 No Cycle
	if CheckCycle(head) == true {
		fmt.Println("has Cycle")
	} else {
		fmt.Println("No Cycle")
	}

	// Test 2.2 Cycle
	CreateCycle(head)
	if CheckCycle(head) == true {
		fmt.Println("has Cycle")
	} else {
		fmt.Println("No Cycle")
	}

	// Test 3 Linked list with More than 2 nodes
	head = nil
	for i := 0; i < 2; i++ {
		node := &Node{Value: i + 1}
		head = InsertNode(node, head)
	}

	// Test 3.1
	if CheckCycle(head) == true {
		fmt.Println("has Cycle")
	} else {
		fmt.Println("No Cycle")
	}

	// Test 3.2 Cycle
	CreateCycle(head)
	if CheckCycle(head) == true {
		fmt.Println("has Cycle")
	} else {
		fmt.Println("No Cycle")
	}

	// Test 4 Linked list with More than 2 nodes
	head = nil
	for i := 0; i <= 2; i++ {
		node := &Node{Value: i + 1}
		head = InsertNode(node, head)
	}
	// Test 4.1
	if CheckCycle(head) == true {
		fmt.Println("has Cycle")
	} else {
		fmt.Println("No Cycle")
	}

	// Test 4.2 Cycle
	CreateCycle(head)
	if CheckCycle(head) == true {
		fmt.Println("has Cycle")
	} else {
		fmt.Println("No Cycle")
	}
}
