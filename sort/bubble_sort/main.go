package main

import (
	"fmt"
)

func BubbleAscending(list []int) {
	l := len(list)
	fmt.Println("len : ", l)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-1-i; j++ {
			fmt.Printf("i : %d, j : %d \n", i, j)
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}

// Let l = 10 (len of array)
// Loop  for I = (0 , 1 ,.....8) (I < L-1)
// Loop for J = (0, 1, 2, 3, 4, 5, 6, 7, 8, 9) (J < L-1-I)
// 	(The index till each point loop continues is L-1-0 . L-2-1, L-3-2, L-4-3, L-5-4, L-6-5, L-7-6, L-8-7, L-9-8)
// 	(9-0 = 9, 9-1 = 9,..., 9-8 = 1) )(L-1 = 9)

// BubbleDescending described
// for a list with len l = 5 , means maxIndex = (l-1) = 4
// outer loop range i := 0 to 3 ( <maxIndex - not less than or equal to)
// inner loop range i := 0 to < (maxIndex-i)(3,2,1,0) ( leave the last element in each loop since it has found its place)
func BubbleDescending(list []int) {
	l := len(list)
	fmt.Println("len : ", l)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-1-i; j++ {
			fmt.Printf("i : %d, j : %d \n", i, j)
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
	// arr := [9]int{10, 91, 65, 23, 14, 67, 89, 54, 78}
	// slice = arr[:]
	// BubbleDescending(slice)
	// fmt.Println(slice)
}
