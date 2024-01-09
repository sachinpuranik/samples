package main

import (
	"fmt"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	var newIntervals [][]int
	var mergedInterval []int

	mergeStart := false
	mergeEnd := false
	for _, interval := range intervals {

		if interval[0] > newInterval[0] {
			mergeStart = true
			mergedInterval = append(mergedInterval, newInterval[0])
		} else {
			mergedInterval = append(mergedInterval, interval[0])
		}

		if !mergeStart && !mergeEnd {
			if interval[1] > newInterval[0] {
				mergeStart = true
				mergedInterval = append(mergedInterval, newInterval[0])
			} else {
				mergedInterval = append(mergedInterval, interval[1])
			}
		} else if mergeStart && !mergeEnd {
			if interval[1] > newInterval[1] {
				mergeStart = false
				mergedInterval = append(mergedInterval, interval[1])
			} else {
				mergedInterval = append(mergedInterval, interval[1])
			}
		}
		newIntervals = append(newIntervals, mergedInterval)
	}
	return newIntervals
}

func main() {
	intervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	newInterval := []int{4, 8}

	fmt.Println(insert(intervals, newInterval))
}
