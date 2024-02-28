package main

import (
	"fmt"
	"unicode/utf8"
)

func strStr(haystack string, needle string) int {
	found := false
	pos := -1
	// n := rune(needle[0])
	n, size := utf8.DecodeRuneInString(needle)
	// n, size := utf8.DecodeRune([]byte(needle))

	fmt.Printf("\n n : %c, size : %d", n, size)

	for i, r := range haystack {
		if r == n {
			end := i + len(needle)
			if end > len(haystack) {
				return pos
			}
			if haystack[i:end] == needle {
				found = true
			}

			if found {
				pos = i
				break
			}
		}
	}
	return pos
}

func main() {
	s := "├this is a string"
	ss := "├th"
	pos := strStr(s, ss)

	fmt.Printf("\n str : %s, find : %s, pos : %d", s, ss, pos)
}
