package main

import (
	"fmt"
)

// QuickSort function to sort an array (with partitioning integrated)
func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	// Choose the last element as pivot
	pivotIndex := len(arr) - 1
	pivot := arr[pivotIndex]

	// Partitioning: rearrange elements around pivot
	i := -1
	// since last element is pivote , we are considering looping only upto LastIndex-1
	for j := 0; j < pivotIndex; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i] // Swap elements smaller than pivot
		}
	}

	// Place the pivot in its correct sorted position
	arr[i+1], arr[pivotIndex] = arr[pivotIndex], arr[i+1]
	// set the new pivot index after swap
	pivotIndex = i + 1

	// Recursively sort left and right subarrays
	QuickSort(arr[:pivotIndex])   // Left subarray
	QuickSort(arr[pivotIndex+1:]) // Right subarray

	return arr
}

func main() {
	fmt.Println(QuickSort([]int{98, 23, 45, 67, 65, 35, 18, 9, 65, 4, 17}))
}
