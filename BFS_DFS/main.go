// ref - https://www.youtube.com/watch?v=PMMc4VsIacU
package main

import (
	"fmt"
)

// Problem - Traversing methods over a graph.
// BFS - breadth first search - Traverse one horizontal level completely before shifting to next level starting from root.
// 		1. This is essentially Queue problem while iterating over nodes. (FIFO nature)
//		2. Each time a node have children add them to queue for processing later.
// 		3. keep looping over queue and de-que the elements and process them until queue is empty.
//		Note : - BFS do not have any concept of Pre , Post order traversal. it's level order traversal.

// DFS - depth first search - Traverse vertically down till you finish all the levels of that sub tree and the shift to another side of subtree.
//		1. This is essentially solved using stack. (LIFO nature)
//		2. Each time a node have children add them to stack for processing later.
// 		3. keep looping over stack. pop the top element (most recent element) and process it. repeat process till the stack is empty
//		Note : it is possible to perform In-oder & Post-order traverse over DFS , but not on BFS

// Node - The two flags marked and visited are important.
// marked - this flag indicate that a node is already covered in the past , so don't try to cover in current cycle.
// visited - flag indicates that we have taken the value and printed or done something about it.
type Node struct {
	Value     int
	Visited   bool
	Marked    bool
	Neighbors Graph
}

type Graph []*Node

func createGraphNodes(n int) Graph {
	graph := make(Graph, 0)
	for i := 0; i < n; i++ {
		graph = append(graph, &Node{
			Value: i,
		})
	}
	return graph
}

func (g Graph) attachChildren(parent int, Neighbors []int) {
	n := g[parent]

	if n.Neighbors == nil {
		n.Neighbors = make(Graph, 0)
	}
	for _, c := range Neighbors {
		n.Neighbors = append(n.Neighbors, g[c])
	}
}

func createBigGraph() Graph {
	graph := createGraphNodes(10)
	graph.attachChildren(0, []int{1, 2})
	graph.attachChildren(1, []int{4, 3})
	graph.attachChildren(2, []int{0, 1})
	graph.attachChildren(3, []int{1, 5})
	graph.attachChildren(4, []int{1})
	graph.attachChildren(5, []int{3, 6, 7, 8})
	graph.attachChildren(6, []int{5})
	graph.attachChildren(7, []int{5, 8})
	graph.attachChildren(8, []int{5, 7, 9})
	graph.attachChildren(9, []int{8})

	return graph
}

func createSmallGraph() Graph {
	graph := createGraphNodes(5)
	graph.attachChildren(0, []int{1, 2, 3})
	graph.attachChildren(1, []int{0, 3})
	graph.attachChildren(2, []int{0, 3})
	graph.attachChildren(3, []int{0, 1, 2, 4})
	graph.attachChildren(4, []int{3})

	return graph
}

func main() {
	graph := createSmallGraph()
	for _, n := range graph {
		fmt.Println("father : %d", n.Value)
		for _, v := range n.Neighbors {
			fmt.Println("	children : %d", v.Value)
		}
	}

	BFSIterative(graph)

	return

	DFSInOrder(graph[0])

	resetGraph(graph)
	DFSIterative(graph)

	resetGraph(graph)
	DFSPostOrder(graph[0])
}

func resetGraph(graph Graph) {
	for _, v := range graph {
		v.Visited = false
		v.Marked = false
	}
}

// DFSInOrder - Visit the node and then children.
func DFSInOrder(n *Node) {
	fmt.Println(n.Value)
	n.Visited = true
	n.Marked = true
	for _, n := range n.Neighbors {
		if n.Marked == false {
			DFSInOrder(n)
		}
	}
}

// DFSPostOrder - Visits the node which have only one neighbour first.
func DFSPostOrder(n *Node) {
	n.Marked = true
	for _, n := range n.Neighbors {
		if n.Marked == false {
			DFSPostOrder(n)
		}
	}
	n.Visited = true
	fmt.Println(n.Value)
}

func DFSIterative(g Graph) {
	fmt.Println("-----DFSIterative Iterative")

	stack := make([]*Node, 0)
	stack = append(stack, g[0])

	for len(stack) > 0 {
		e := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !e.Marked {
			e.Visited = true
			fmt.Println(e.Value)
			e.Marked = true

			for _, n := range e.Neighbors {
				if !n.Marked {
					stack = append(stack, n)
				}
			}
		}

	}
}

func DFSIterativePostorder(g Graph) {
	fmt.Println("-----DFSIterative Postorder")

	stack := make([]*Node, 0)  // First stack for DFS
	output := make([]*Node, 0) // Second stack for Postorder

	stack = append(stack, g[0]) // Start with the root node

	for len(stack) > 0 {
		// Pop from the stack
		e := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Push the node to the output stack
		output = append(output, e)

		// Add neighbors to the stack (visit later)
		for _, n := range e.Neighbors {
			if !n.Marked {
				stack = append(stack, n)
			}
		}
	}

	// Process nodes in reverse order from the output stack
	for i := len(output) - 1; i >= 0; i-- {
		node := output[i]
		if !node.Marked {
			fmt.Println(node.Value)
			node.Marked = true
		}
	}
}

// BFSIterative - this also does level order traversal.
// means one level of nodes are processed at first, once they are done , you shift to next level down.
// this achieved through queue.
// you visit the root , and add its children to a queue ,
// start processing first child after root is done. if it has children add it to queue.
// loop until queue is empty.
func BFSIterative(g Graph) {
	fmt.Println("-----BFSIterative Iterative")

	q := make([]*Node, 0)
	q = append(q, g[0])

	for len(q) > 0 {
		e := q[0]
		q = q[1:]

		if !e.Marked {
			e.Visited = true
			fmt.Println(e.Value)
			e.Marked = true

			for _, n := range e.Neighbors {
				if !n.Marked {
					q = append(q, n)
				}
			}
		}
	}
}
