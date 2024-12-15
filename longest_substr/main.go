package main

import (
	"fmt"
)

// Problem - def
// Find the lomgest substr from a given string

func getLen(s string) int {
	counter := make(map[rune]int, 0)
	for _, r := range s {
		if _, found := counter[r]; found == true {
			break
		}
		counter[r] = 1
	}
	return len(counter)
}

func lengthOfLongestSubstring(s string) int {
	maxLen := 5 * 10 * 10 * 10 * 10
	if len(s) > maxLen {
		return -1
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

// Function to efficiently find the longest substring without repeating characters
func lengthLongestUniqueSubstringGemini(s string) int {
	var startIdx, endIdx int
	start, end, maxLength := 0, 0, 0
	seen := make(map[byte]int) // Use a map to track seen characters efficiently

	for i := 0; i < len(s); i++ {
		// Check if the current character is already seen
		ss := s[i]
		if pos, ok := seen[ss]; ok {
			// Update start index if the previous occurrence is beyond the current substring
			if pos > start {
				start = pos + 1
			}
		}

		// Update the seen map and update potential max length
		seen[s[i]] = i
		if end-start+1 > maxLength {
			maxLength = end - start + 1
			// Update potential substring based on updated start and max length
			startIdx, endIdx = start, end+1 // +1 to include the current character
		}

		end = i // Update end of the current potential substring
	}

	// Return the longest substring
	return len(s[startIdx:endIdx])
}

// This is calculating pure length and not the substring
func lengthOfLongestSubstringGPT(s string) int {
	charIndex := make(map[rune]int)
	maxLength := 0
	startIndex := 0
	preservedStartIndex := 0

	for i, char := range s {
		index, found := charIndex[char]
		if found && index >= startIndex {
			startIndex = index + 1
		}

		charIndex[char] = i
		currentLength := i - startIndex + 1
		if currentLength > maxLength {
			preservedStartIndex = startIndex
			maxLength = currentLength
		}
		fmt.Println(charIndex)
	}

	fmt.Println(charIndex)
	fmt.Println(s[preservedStartIndex : preservedStartIndex+maxLength])

	return maxLength
}

func main() {
	// expect 3,1,3,14
	set := []string{"ababcdaxyzabc", "abcabcbb", "bbbbb", "pwwkew", "xhhyccrcbdczkvzeeubynglxfdedshtpobqsdhufkzgwuhaabdzrlkosnuxibrxssnkxuhcggkecshdvkcmymdqbxolbfjtzyfw"}
	for _, s := range set {
		fmt.Printf("\n len %s:(%d)", s, lengthOfLongestSubstringGPT(s))
	}
}
