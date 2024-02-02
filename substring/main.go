package main

import (
	"errors"
	"fmt"
)

func match(inString string, findStr string, startIndex int) bool {
	for i := 0; i < len(findStr); i++ {
		if findStr[i] != inString[startIndex+i] {
			return false
		}
	}
	return true
}

// CheckSubstring -
func CheckSubstring(inString, findStr string) (int, error) {
	if len(inString) == 0 || len(findStr) == 0 {
		return -1, errors.New("Blank input. Check the input")
	}
	for index := range inString {
		if inString[index] == findStr[0] {
			if len(findStr) > len(inString[index:]) {
				return -1, nil
			}
			if match(inString, findStr, index) == true {
				return index, nil
			}
		}
	}
	return -1, nil
}

func test1() {
	s1 := "My name is sachin"
	s2 := "sac"
	location, err := CheckSubstring(s1, s2)

	if err != nil {
		fmt.Println("Some error finding the string")
	} else if location == -1 {
		fmt.Println("string not found")
	} else {
		fmt.Println("string found at location =>", location)
	}
}

func test2() {
	s1 := ""
	s2 := "sac"
	location, err := CheckSubstring(s1, s2)

	if err != nil {
		fmt.Println("Error =>", err.Error())
	} else if location == -1 {
		fmt.Println("string not found")
	} else {
		fmt.Println("string found at location =>", location)
	}
}

func test3() {
	s1 := "My name is sachin"
	s2 := ""
	location, err := CheckSubstring(s1, s2)

	if err != nil {
		fmt.Println("Error =>", err.Error())
	} else if location == -1 {
		fmt.Println("string not found")
	} else {
		fmt.Println("string found at location =>", location)
	}
}

func test4() {
	s1 := "sa"
	s2 := "sac"
	location, err := CheckSubstring(s1, s2)

	if err != nil {
		fmt.Println("Error =>", err.Error())
	} else if location == -1 {
		fmt.Println("string not found")
	} else {
		fmt.Println("string found at location =>", location)
	}
}

func test5() {
	s1 := "My name is sachin"
	s2 := "hwa"
	location, err := CheckSubstring(s1, s2)

	if err != nil {
		fmt.Println("Error =>", err.Error())
	} else if location == -1 {
		fmt.Println("string not found")
	} else {
		fmt.Println("string found at location =>", location)
	}
}

func main() {
	test1()
	test2()
	test3()
	test4()
	test5()
}
