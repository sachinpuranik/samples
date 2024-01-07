package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	maxLen := 5 * 10 * 10 * 10 * 10
	if len(s) > maxLen {
		return -1
	}
	getLen := func(s string) int {
		counter := make(map[rune]int, 0)
		for i := 0; i < len(s); i++ {
			r := rune(s[i])
			if _, found := counter[r]; found == true {
				break
			}
			counter[r] = 1
		}
		return len(counter)
	}

	finalLen := 0
	subLen := 0
	for i := 0; i < len(s); i++ {
		subStr := s[i:]
		// fmt.Println(subStr)
		if len(subStr) <= finalLen {
			break
		}
		subLen = 0
		for j := 0; j < len(subStr); j++ {
			s := subStr[j:]
			// fmt.Println(s)
			if len(s) <= finalLen {
				break
			}
			l := getLen(s)
			if subLen < l {
				subLen = l
			}
		}
		if finalLen < subLen {
			finalLen = subLen
		}
		if finalLen >= 94 {
			break
		}
	}
	return finalLen
}

func main() {
	// expect 3,1,3,14
	set := []string{"abcabcbb", "bbbbb", "pwwkew", "xhhyccrcbdczkvzeeubynglxfdedshtpobqsdhufkzgwuhaabdzrlkosnuxibrxssnkxuhcggkecshdvkcmymdqbxolbfjtzyfw"}
	for _, s := range set {
		fmt.Printf("\n len %s:(%d)", s, lengthOfLongestSubstring(s))
	}
}
