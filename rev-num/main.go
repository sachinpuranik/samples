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

		// this check ensures number does not overflow out of max int size
		// next rev num is calculated by multiplying by 10 and adding the poped digit. so we check exactly reverse of this.
		// if current rev number is already bigger than calculated max number by calculation (maxInt32-pop)/10 the, after performing the next operation it will overflow.
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
