// ref - https://www.youtube.com/watch?v=PMMc4VsIacU
package main

import (
	"fmt"
	// "math/rand/v2"
)

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

	DFSInOrder(graph, graph[0])

	resetGraph(graph)
	DFSIterative(graph)

	resetGraph(graph)
	DFSPostOrder(graph, graph[0])
}

func resetGraph(graph Graph) {
	for _, v := range graph {
		v.Visited = false
		v.Marked = false
	}
}

// DFSInOrder - Visit the node and then children.
func DFSInOrder(g Graph, n *Node) {
	fmt.Println(n.Value)
	n.Visited = true
	n.Marked = true
	for _, n := range n.Neighbors {
		if n.Marked == false {
			DFSInOrder(g, n)
		}
	}
}

// DFSPostOrder - Visits the node which have only one neighbour first.
func DFSPostOrder(g Graph, n *Node) {
	n.Marked = true
	for _, n := range n.Neighbors {
		if n.Marked == false {
			DFSPostOrder(g, n)
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
