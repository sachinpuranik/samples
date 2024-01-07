package main

import (
	"fmt"
)

func bubbleSort(arr []int) {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-1-i; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{10, 90, 65, 23, 11, 67, 89, 54, 65}
	bubbleSort(arr)
	fmt.Println(arr)
}
