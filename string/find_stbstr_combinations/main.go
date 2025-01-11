// Example 1:

// Input: s1 = "ab", s2 = "eidbaooo"
// Output: true
// Explanation: s2 contains one permutation of s1 ("ba").
// Example 2:

// Input: s1 = "ab", s2 = "eidboaoo"
// Output: false

// https://leetcode.com/problems/permutation-in-string/
// important note , the letters can repeat in input or in the substr. so its important to calculate heat map for each string or substring.
// createHash function calculates the heatmap for a string

package main

import (
	"fmt"
)

func compareHash(m1 map[rune]int, m2 map[rune]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v := range m1 {
		count, found := m2[k]
		if !found || count != v {
			return false
		}
	}
	return true
}

func createHash(s string) map[rune]int {
	hash := make(map[rune]int)
	for _, v := range s {
		hash[v]++
	}
	return hash
}

func findSubstrCombination(inputStr string, subStr string) bool {
	if len(inputStr) == 0 {
		return false
	}

	hSubStr := createHash(subStr)

	for index, v := range inputStr {
		_, found := hSubStr[v]
		if found && index+len(subStr) <= len(inputStr) {
			inputHash := createHash(inputStr[index : index+len(subStr)])
			res := compareHash(hSubStr, inputHash)
			if res {
				return res
			}
		}
	}

	return false
}

// computeMap
func computeMap(s string) map[string]bool {
	var stringHash map[string]bool
	stringHash = make(map[string]bool)
	for _, r := range s {
		stringHash[string(r)] = true
	}
	return stringHash
}

func setup(str, subStr string) {
	found := findSubstrCombination(str, subStr)
	if found == true {
		fmt.Println("Combination found")
	} else if found == false {
		fmt.Println("Combination not found")
	}
}

func test1() {
	s1 := "abcdefaa"
	s2 := "faa"
	setup(s1, s2)
}

func test2() {
	s1 := "abcdefaa"
	s2 := "aaf"
	setup(s1, s2)
}

func test3() {
	s1 := "abcdefaa"
	s2 := "afa"
	setup(s1, s2)
}

func test4() {
	s1 := "abcdefaa"
	s2 := "afaz"
	setup(s1, s2)
}

func main() {
	test1()
	test2()
	test3()
	test4()
}
