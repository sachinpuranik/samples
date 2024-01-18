package main

import (
	"fmt"
)

func BubbleAscending(list []int) {
	l := len(list)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-1-i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}

// BubbleDescending described
// for a list with len l = 5 , means maxIndex = (l-1)
// outer loop range i := 0 to 3 ( <maxIndex - not less than or equal to)
// inner loop range i := 0 to < maxIndex-i ( leave the last element in each loop since it has found its place)
func BubbleDescending(list []int) {
	l := len(list)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-1-i; j++ {
			if list[j] < list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}

func main() {
	// caution - arr here is a slice and not array. So you can pass a slice to a function as and argument and can be changed. it will reflect.
	slice := []int{10, 90, 65, 23, 11, 67, 89, 54, 65}
	BubbleAscending(slice)
	fmt.Println(slice)

	// if you have to pass an array , you need to convert it to slice and then pass.
	arr := [9]int{10, 91, 65, 23, 14, 67, 89, 54, 78}
	slice = arr[:]
	BubbleDescending(slice)
	fmt.Println(slice)
}
