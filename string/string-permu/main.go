package main

import (
	"fmt"
)

// given input - 123 - program prints as below.
// 123
// 132
// 213
// 231
// 312
// 321
// 12
// 13
// 21
// 23
// 31
// 32
// 1
// 2
// 3

func substring(content string, part string, index int) {
	if index >= len(content) {
		return
	}
	sub := part + content[index:index+1]

	substring(content, sub, index+1)

	permutation("", sub)
}

func permutation(prefix, str string) {
	n := len(str)
	if n == 0 {
		fmt.Println(prefix)
	} else {
		for i := 0; i < n; i++ {
			permutation(prefix+str[i:i+1], str[0:i]+str[i+1:n])
		}
	}
}

func generatePermutations(s string) []string {
	result := []string{}

	if len(s) == 1 {
		result = append(result, s)
		return result
	}
	for i, rune := range s {
		subStr := s[0:i] + s[i+1:]
		pastResult := generatePermutations(subStr)
		for _, past := range pastResult {
			result = append(result, string(rune)+past)
		}
		result = append(result, string(rune))
	}
	return result
}

func main() {
	s := "1234"

	result := generatePermutations(s)
	for _, r := range result {
		fmt.Println(r)
	}
}
