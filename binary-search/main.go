package main

import (
	"fmt"
)

func BinarySearchCustom(list []int, n int) int {
	i := 0
	j := len(list) - 1

	if list[i] == n {
		return i
	}

	if list[j] == n {
		return j
	}

	for i < j {
		mid := (i + j) / 2
		if mid == i || mid == j {
			return -1
		}
		if list[mid] == n {
			return mid
		} else if n < list[mid] {
			j = mid
		} else if n > list[mid] {
			i = mid
		}
	}
	return -1
}

// BinarySearch performs binary search on a sorted array.
func BinarySearchOfficial(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid // Target found, return its index
		} else if arr[mid] < target {
			low = mid + 1 // Target is in the right half
		} else {
			high = mid - 1 // Target is in the left half
		}
	}

	return -1 // Target not found
}

func main() {
	list := []int{11, 14, 17, 21, 24, 27, 32, 35, 39, 42, 47, 52}
	fmt.Printf("list : %v , found-pos : %d", list, BinarySearchCustom(list, 52))
}
