package main

import (
	"fmt"
)

// Algo
// 1. If new interval is completely bigger than the one in the list, append the element from list to result.
// 2. If new interval is completely smaller than the one in the list, append new interval in the list. and now interval becomes merge-interval. Candidate for next merge.
// 3. If the intervals overlap , take minimum of 0th element , and mx of 1st element of both intervals and form a newInterval. - This is basically merge step.
// 4. Finally append the pending merge-interval as a last element to result and return

func insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	mergedInterval := newInterval

	for _, interval := range intervals {
		if interval[1] < mergedInterval[0] {
			// Non-overlapping, before new interval
			result = append(result, interval)
		} else if interval[0] > mergedInterval[1] {
			// Non-overlapping, after new interval
			result = append(result, mergedInterval)
			mergedInterval = interval
		} else {
			// Overlapping, merge intervals
			mergedInterval[0] = min(mergedInterval[0], interval[0])
			mergedInterval[1] = max(mergedInterval[1], interval[1])
		}
	}

	result = append(result, mergedInterval)
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	intervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	newInterval := []int{4, 8}

	fmt.Println(insert(intervals, newInterval))
}
