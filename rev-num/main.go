package main

import (
	"fmt"
	"math"
)

func reverseNumber(x int) int {
	maxInt32, _ := math.MaxInt32, math.MinInt32

	n := x

	revN := 0
	sign := 1

	if n < 0 {
		sign = -1 // take out the sign
		n = -n    // negate the number
	}

	for n > 0 {

		pop := n % 10
		n = n / 10

		if revN > (maxInt32-pop)/10 {
			revN = -1
			break
		}

		revN = revN*10 + pop

	}
	revN = sign * revN

	return revN
}

func main() {
	x := -1234
	rx := reverseNumber(x)
	fmt.Printf("\n number %d , reverse number : %d", x, rx)
}
