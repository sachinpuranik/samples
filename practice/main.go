package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	next *Node
	val  int
}

type DNode struct {
	next *DNode
	prev *DNode
	val  int
}

func Insert(n *Node, head *Node) *Node {
	n.next = head
	head = n
	return head
}

func dInsert(n *DNode, head *DNode) *DNode {
	if head == nil {
		n.next = nil
		n.prev = nil
		head = n
		return head
	}
	n.next = head
	head.prev = n
	head = n
	return head
}

func createSLL() *Node {
	var head *Node

	n := &Node{nil, 10}
	head = Insert(n, head)

	n = &Node{nil, 20}
	head = Insert(n, head)

	n = &Node{nil, 30}
	head = Insert(n, head)

	n = &Node{nil, 40}
	head = Insert(n, head)

	n = &Node{nil, 50}
	head = Insert(n, head)

	return head
}

func createCycle(head *Node) *Node {
	tmp := head
	if tmp != nil {
		for tmp.next != nil {
			tmp = tmp.next
		}
	}
	tmp.next = head
	return head
}

// basic algo as below.
// 0. first step if lenghth of list is 0 or one , loop can not exits.
// 1. two pointers slow and fast walk behind each other.
// 2. slow walks 1 step vs fast walk 2 steps always.
// 3. means fast will always approach the last node fast , else it will reach slow , it loop exist.
// 4. loop when fast is not null , or not equal to slow. (these are only two breaking conditions)
//  5. check if slow and fast are same, then break as loop exist.
//  6. check if next of fast is not nil and next.next  is not nil either. Now its safe to move two positions ahead on list.
//  7. since we already ensured that fast pointer walks smooth , slow will be save to move ahead. so slow moves slow.next.
//
// Continue through 4-7 till loop is found or pointer reaches nil. and you know the result.
func detectCycle(head *Node) bool {
	// Handle empty or single-node lists directly:
	if head == nil || head.next == nil {
		return false
	}

	// Use descriptive variable names:
	slowPointer := head
	fastPointer := head.next

	// Loop until fastPointer reaches the end or a cycle is detected:
	for fastPointer != nil && fastPointer.next != nil {
		if fastPointer == slowPointer {
			return true // Cycle detected
		}

		// Move pointers at different speeds:
		slowPointer = slowPointer.next
		fastPointer = fastPointer.next.next
	}

	return false
}

func createDLL() *DNode {
	var head *DNode
	head = dInsert(&DNode{nil, nil, 10}, head)
	head = dInsert(&DNode{nil, nil, 20}, head)
	head = dInsert(&DNode{nil, nil, 30}, head)
	head = dInsert(&DNode{nil, nil, 40}, head)
	return head
}

func printSLL(head *Node) {
	t := head
	fmt.Println("Prinitng SLL")
	for t != nil {
		fmt.Println(t.val)
		t = t.next
	}
}

func printDLL(head *DNode) {
	t := head
	fmt.Println("Prinitng DLL")
	for t.next != nil {
		fmt.Println(t.val)
		t = t.next
	}

	for t != head {
		fmt.Println(t.val)
		t = t.prev
	}

	fmt.Println(t.val)
}

// Conditions -  No Repeat, possibly take from 0 , makes it easy.

func TwoSum(target int) (index1 int, index2 int) {
	input := []int{1, 4, 10, 6, 9, 2, 3, 14, 5, 7, 12, 11, 8, 0}
	var ih map[int]int
	ih = make(map[int]int, 1)
	for i, val := range input {
		ih[val] = i
	}

	for i, val := range input {
		n2 := target - val
		i2, ok := ih[n2]
		if ok {
			return i, i2
		}
	}

	return -1, -1
}

func ReverseList(head *Node) *Node {
	var newHead, tmp *Node

	tmp = head

	for head != nil {
		head = head.next
		tmp.next = newHead
		newHead = tmp
		tmp = head
	}

	for tmp != nil {
		tmp.next = newHead
		newHead = tmp
	}

	return newHead
}

// // ajbsoa
func LongestSubstr(s string) (int, string) {
	tracker := map[rune]int{}
	startIndex := 0
	length := 0
	startIndexLongest := 0

	for i, val := range s {
		oldIndex, found := tracker[val]
		if found && oldIndex >= startIndex {
			startIndex = oldIndex + 1
		}
		tracker[val] = i
		newLength := i - startIndex + 1
		if newLength > length {
			startIndexLongest = startIndex
			length = newLength
		}

	}

	subs := s[startIndexLongest : startIndexLongest+length]

	return length, subs
}

func Bubble(in []int) {
	l := len(in)

	for i := 0; i < l-1; i++ {
		for j := 0; j < l-1-i; j++ {
			if in[j] > in[j+1] {
				in[j], in[j+1] = in[j+1], in[j]
			}
		}
	}

	fmt.Println(in)
}

func BinarySearch(n int, in []int) bool {
	low := 0
	high := len(in)
	mid := low

	for low <= high {
		mid = low + (high-low)/2

		if in[mid] == n {
			return true
		} else if in[mid] < n {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}

// -1,0,1,error
func checkVersion(v1 string, v2 string) (int, error) {
	detailedV1 := strings.Split(v1, ".")
	detailedV2 := strings.Split(v2, ".")

	l1 := len(detailedV1)
	l2 := len(detailedV2)

	if l1 < l2 {
		for i := l1; i < l2; i++ {
			detailedV1 = append(detailedV1, "0")
		}
	} else if l2 < l1 {
		for i := l2; i < l1; i++ {
			detailedV2 = append(detailedV2, "0")
		}
	}

	for i, v1s := range detailedV1 {
		v2s := detailedV2[i]

		nv1, err := strconv.Atoi(v1s)
		if err != nil {
			return -1, err
		}

		nv2, err := strconv.Atoi(v2s)
		if err != nil {
			return -1, err
		}

		if nv1 > nv2 {
			return -1, nil
		} else if nv2 > nv1 {
			return 1, nil
		}
	}
	return 0, nil
}

func Merge(left, right []int) []int {
	merge := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			merge = append(merge, right...)
			return merge
		} else if len(right) == 0 {
			merge = append(merge, left...)
			return merge
		} else if left[0] < right[0] {
			merge = append(merge, left[0])
			left = left[1:]
		} else {
			merge = append(merge, right[0])
			right = right[1:]
		}
	}

	return merge
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return Merge(left, right)
}

type TNode struct {
	left  *TNode
	right *TNode
	val   string
}

func traverseLRR(root *TNode) {
	if root.left != nil {
		traverseLRR(root.left)
	}
	fmt.Println(root.val)
	if root.right != nil {
		traverseLRR(root.right)
	}
}

func traverseRLR(root *TNode) {
	fmt.Println(root.val)
	if root.left != nil {
		traverseRLR(root.left)
	}
	if root.right != nil {
		traverseRLR(root.right)
	}
}

func traverseRRL(root *TNode) {
	fmt.Println(root.val)
	if root.right != nil {
		traverseRRL(root.right)
	}
	if root.left != nil {
		traverseRRL(root.left)
	}
}

func createBinaryTree() *TNode {
	var root *TNode

	left := &TNode{&TNode{nil, nil, "a"}, &TNode{nil, nil, "b"}, "i"}
	right := &TNode{&TNode{nil, nil, "c"}, &TNode{nil, nil, "d"}, "j"}

	t1 := &TNode{left, right, "m"}

	left = &TNode{&TNode{nil, nil, "e"}, &TNode{nil, nil, "f"}, "k"}
	right = &TNode{&TNode{nil, nil, "g"}, &TNode{nil, nil, "h"}, "l"}

	t2 := &TNode{left, right, "n"}

	root = &TNode{t1, t2, "o"}

	return root
}

func main() {
	// sll := createSLL()
	// printSLL(sll)
	// sll = createCycle(sll)
	// fmt.Println("cycle => ", detectCycle(sll))

	// sll = ReverseList(sll)
	// printSLL(sll)

	// dll := createDLL()
	// printDLL(dll)

	// i1, i2 := TwoSum(50)
	// println(i1, i2)

	// fmt.Println(LongestSubstr("abcdac"))

	// fmt.Println(checkVersion("11.11.13", "11.11.12"))

	// fmt.Println(LongestSubstr("abcdcbzxy"))

	// Sorts
	// arr := []int{9, 2, 56, 68, 52, 75, 39, 64, 12, 34, 54, 73}
	// Bubble(arr)
	// fmt.Println(MergeSort(arr))

	// r := createBinaryTree()
	// traverseLRR(r)
	// traverseRLR(r)
	// traverseRRL(r)

	// fmt.Println(findSubstrCombination("abcdefaa", "aafe"))
	// printAllSubstrings("abcd")

	TestGraph()
}

func compareHash(m1 map[rune]int, m2 map[rune]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v := range m1 {
		count, found := m2[k]
		if !found || count != v {
			return false
		}
	}
	return true
}

func createHash(s string) map[rune]int {
	hash := make(map[rune]int)
	for _, v := range s {
		hash[v]++
	}
	return hash
}

func findSubstrCombination(inputStr string, subStr string) bool {
	if len(inputStr) == 0 {
		return false
	}

	hSubStr := createHash(subStr)

	for index, v := range inputStr {
		_, found := hSubStr[v]
		if found && index+len(subStr) <= len(inputStr) {
			inputHash := createHash(inputStr[index : index+len(subStr)])
			res := compareHash(hSubStr, inputHash)
			if res {
				return res
			}
		}
	}

	return false
}

func printAllSubstrings(input string) {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j <= len(input); j++ { // Start j from i+1 and keep j <= len(input)
			substr := input[i:j]
			fmt.Println(substr)
		}
	}
}

type GNode struct {
	mark      bool
	visited   bool
	value     int
	neighbors graph
}

type graph []*GNode

func TestGraph() {
	g := createGraph()
	DFSTraverse(g)

	g = createGraph()
	BFSTraverse(g)
}

func createLinks(g graph, node *GNode, neighbors []int) graph {
	for _, v := range neighbors {
		node.neighbors = append(node.neighbors, g[v])
	}
	return g
}

func createGraphNodes(nodeCount int) graph {
	var g graph
	for i := 0; i < nodeCount; i++ {
		n := &GNode{value: i}
		g = append(g, n)
	}

	return g
}

func createGraph() graph {
	g := createGraphNodes(10)

	g = createLinks(g, g[0], []int{2, 3})
	g = createLinks(g, g[2], []int{5, 7})
	g = createLinks(g, g[3], []int{4, 6, 1})
	g = createLinks(g, g[6], []int{9, 8})
	g = createLinks(g, g[8], []int{1, 0})

	return g
}

func DFSTraverse(g graph) {
	fmt.Println("Running DFS")

	stack := make(graph, 0, 0)
	stack = append(stack, g[0])

	for len(stack) > 0 {
		// pop element
		e := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		e.mark = true
		e.visited = true
		fmt.Println(e.value)

		if e.neighbors != nil {
			for _, child := range e.neighbors {
				if !child.mark {
					stack = append(stack, child)
				}
			}
		}

	}
}

func BFSTraverse(g graph) {
	fmt.Println("Running BFS")

	queue := make(graph, 0, 0)
	queue = append(queue, g[0])

	for len(queue) > 0 {
		// pop element
		e := queue[0]
		queue = queue[1:]

		e.mark = true
		e.visited = true
		fmt.Println(e.value)

		if e.neighbors != nil {
			for _, child := range e.neighbors {
				if !child.mark {
					queue = append(queue, child)
				}
			}
		}

	}
}
