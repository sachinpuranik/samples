package main

import (
	"fmt"
)

// Node -
type Node struct {
	Next  *Node
	Value int
}

// CheckCycleCustom -
func CheckCycleCustom(head *Node) bool {
	if head == nil || head.Next == nil {
		return false //
	}

	follow, runner := head, head.Next

	for runner != nil || runner != follow {
		if runner == follow {
			return true
		}
		if runner.Next != nil && runner.Next.Next != nil {
			runner = runner.Next.Next
		} else {
			return false
		}
		follow = follow.Next
	}
	return false
}

// CheckCycleSimplified
func CheckCycleSimplified(head *Node) bool {
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
	if CheckCycleSimplified(head) == true {
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
	if CheckCycleSimplified(head) == true {
		fmt.Println("has Cycle")
	} else {
		fmt.Println("No Cycle")
	}

	// Test 2.2 Cycle
	CreateCycle(head)
	if CheckCycleSimplified(head) == true {
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
	if CheckCycleSimplified(head) == true {
		fmt.Println("has Cycle")
	} else {
		fmt.Println("No Cycle")
	}

	// Test 3.2 Cycle
	CreateCycle(head)
	if CheckCycleSimplified(head) == true {
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
	if CheckCycleSimplified(head) == true {
		fmt.Println("has Cycle")
	} else {
		fmt.Println("No Cycle")
	}

	// Test 4.2 Cycle
	CreateCycle(head)
	if CheckCycleSimplified(head) == true {
		fmt.Println("has Cycle")
	} else {
		fmt.Println("No Cycle")
	}
}
