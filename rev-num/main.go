package main

import (
	"fmt"
	"math"
)

func main() {
	maxInt32, _ := math.MaxInt32, math.MinInt32

	x := -1234

	n := x

	revN := 0
	sign := 1

	if n < 0 {
		sign = -1
		n = -n
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

	fmt.Printf("\n number %d , reverse number : %d", x, revN)
}
