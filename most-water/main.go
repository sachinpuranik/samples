package main

// ref : https://leetcode.com/problems/container-with-most-water/?envType=list&envId=du693s
import (
	"fmt"
)

func maxAreaSlow(height []int) int {
	maxVolume := 0

	calculateVolume := func(bars []int) int {
		maxHeight := bars[0]
		if bars[0] > bars[len(bars)-1] {
			maxHeight = bars[len(bars)-1]
		}
		volume := maxHeight * (len(bars) - 1)
		return volume
	}

	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			vol := calculateVolume(height[i : j+1])
			if vol > maxVolume {
				maxVolume = vol
			}
		}
	}
	return maxVolume
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func maxAreaFast(height []int) int {
	maxVolume := 0
	l := 0
	r := len(height) - 1

	for r > l {

		maxHeight := max(height[l], height[r])

		vol := maxHeight * (r - l)

		if vol > maxVolume {
			maxVolume = vol
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxVolume
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	area := maxAreaFast(height)

	if area == 49 {
		fmt.Println("pass", area)
	} else {
		fmt.Println("fail", area)
	}

	height = []int{1, 9, 6, 2, 5, 4, 8, 9, 7}

	area = maxAreaFast(height)

	if area == 54 {
		fmt.Println("pass", area)
	} else {
		fmt.Println("fail", area)
	}
}
