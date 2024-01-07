// Example 1:

// Input: s1 = "ab", s2 = "eidbaooo"
// Output: true
// Explanation: s2 contains one permutation of s1 ("ba").
// Example 2:

// Input: s1 = "ab", s2 = "eidboaoo"
// Output: false

package main

import (
	"errors"
	"fmt"
)

func matchFound(searchIn string, index int, phrase map[string]bool) bool {
	if len(phrase)+index > len(searchIn) {
		return false
	}

	for i := index; i < (len(phrase) + index); i++ {
		_, ok := phrase[searchIn[i:i+1]]
		if ok != true {
			return false
		}
	}
	return true
}

// checkPermutation -
func checkPermutation(searchIn string, phrase map[string]bool) (bool, error) {
	if len(searchIn) == 0 && len(phrase) == 0 {
		return false, errors.New("Input paran len can not be 0")
	}

	if len(searchIn) < len(phrase) {
		return false, errors.New("Search string should be smaller than input string")
	}

	for index := range searchIn {
		_, ok := phrase[searchIn[index:index+1]]
		if ok == true {
			if matchFound(searchIn, index, phrase) == true {
				return true, nil
			}
		}
	}

	return false, nil
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
	hmap := computeMap(subStr)
	found, err := checkPermutation(str, hmap)
	if err != nil {
		fmt.Println("Error =>", err.Error())
	} else if found == true {
		fmt.Println("Combination found")
	} else if found == false {
		fmt.Println("Combination not found")
	}
}

func test1() {
	s1 := "my name is sachin"
	s2 := "ym"
	setup(s1, s2)
}

func test2() {
	s1 := "my name is sachin"
	s2 := "na"
	setup(s1, s2)
}

func main() {
	test1()
	test2()
}
