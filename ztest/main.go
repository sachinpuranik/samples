package main

import (
	"fmt"
)

func main() {
	s := "this is a string"
	l := len(s)
	fmt.Printf("len : %d , substr : %v ", l, s[1:l])

	n := 123
	fmt.Printf("\n mod and div %d %d", n%10, n/10)
}
