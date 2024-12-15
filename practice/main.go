package main

import (
	"fmt"
	"strconv"
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

	n := &Node{
		nil, 10,
	}

	head = Insert(n, head)
	n = &Node{
		nil, 20,
	}
	head = Insert(n, head)
	n = &Node{
		nil, 30,
	}
	head = Insert(n, head)
	n = &Node{
		nil, 40,
	}
	head = Insert(n, head)
	n = &Node{
		nil, 50,
	}

	return head
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

func main() {
	// sll := createSLL()
	// printSLL(sll)

	// sll = ReverseList(sll)
	// printSLL(sll)

	// dll := createDLL()
	// printDLL(dll)

	// i1, i2 := TwoSum(50)
	// println(i1, i2)

	// fmt.Println(LongestSubstr("abcdac"))

	// Bubble([]int{9, 2, 56, 68, 52, 75, 39, 64, 12, 34, 54, 73})

	fmt.Println(checkVersion("11.11", "11.12"))
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
// func LongestSubstr(s string) (int, string) {
// 	hold := map[rune]int{}
// 	prevStart := 0
// 	MaxLen:=

// 	for i, val := range s {
// 		oldIndex, found := hold[val]
// 		if found {
// 			prevStart = oldIndex
// 			l = 0
// 		}
// 		hold[val] = i
// 		l++
// 	}

// 	sub := s[startIndex : startIndex+l]

// 	return l, sub
// }

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
	l1 := len(v1)
	l2 := len(v2)

	if l1 < l2 {
		for i := l1; i <= l2; i++ {
			v1 = v1 + "0"
		}
	} else if l2 < l1 {
		for i := l2; i <= l1; i++ {
			v2 = v2 + "0"
		}
	}

	nv1, err := strconv.Atoi(v1)
	if err != nil {
		return -1, err
	}

	nv2, err := strconv.Atoi(v1)
	if err != nil {
		return -1, err
	}

	if nv1 > nv2 {
		return -1, nil
	} else if nv2 > nv1 {
		return 1, nil
	}

	return 0, nil
}
