package main

import (
	"fmt"
)

// searchMatrixSimple , first find the row in which the element may be present , then find the actual number with binary search.
func searchMatrixSimple(metrics [][]int, num int) bool {
	n := len(metrics[0])

	result := false
	for _, row := range metrics {
		if num == row[n-1] {
			return true
		} else if num < row[n-1] {
			i := 0
			j := n - 1
			for i <= j {
				mid := (i + j) / 2
				if num == row[mid] {
					return true
				} else if num < row[mid] {
					j = mid - 1
				} else if num > row[mid] {
					i = mid + 1
				}

			}
		}
	}

	return result
}

// searchMatrixAdvanced - loop starting from 0th row and max col ,
// increase row in each loop when you know that the number is in next row.
// once you find the relavent row , start decreasing the column , to find the element. (it almost loops through all the elements)
func searchMatrixAdvanced(matrix [][]int, n int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rows, cols := len(matrix), len(matrix[0])
	row, col := 0, cols-1

	for row < rows && col >= 0 {
		if matrix[row][col] == n {
			return true
		} else if matrix[row][col] > n {
			col--
		} else {
			row++
		}
	}

	return false
}

func main() {
	m := [][]int{
		{1, 4, 6, 8},
		{14, 20, 25, 28},
		{31, 38, 42, 46},
		{48, 52, 56, 58},
	}

	fmt.Println("found : ", searchMatrixSimple(m, 45))
}
