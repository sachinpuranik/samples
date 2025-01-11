package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Assumptions - Version no can have multiple "." in the string.
// Core algo
// 1. Split both the version numbers at each "." and form a array of each substring
// 2. Once you have arrays , check which array is small in length. Add the preceding 0's to the smaller array , just to make them exactly same in length
// 3. Now start looping on array from 0th position.
// 	4. for each index , convert both the strings to integer , if error , return error , else compare the numbers and return the response accordingly.

// CompareVersion -
func CompareVersion(v1str, v2str string) (int, error) {
	// return -1 if v1<v2, 1 if v1 > v2 , 0 if v1 == v2
	v1 := strings.Split(v1str, ".")
	v2 := strings.Split(v2str, ".")

	if len(v2) > len(v1) {
		for index := len(v1); index < len(v2); index++ {
			v1 = append(v1, "0")
		}
	} else if len(v1) > len(v2) {
		for index := len(v2); index < len(v1); index++ {
			v2 = append(v2, "0")
		}
	}

	for index := 0; index < len(v1); index++ {
		v1int, err1 := strconv.Atoi(v1[index])
		if err1 != nil {
			return 0, errors.New("Improper v1 string ")
		}

		v2int, err2 := strconv.Atoi(v2[index])
		if err2 != nil {
			return 0, errors.New("Improper v2 string ")
		}

		if v1int > v2int {
			return -1, nil
		} else if v2int > v1int {
			return 1, nil
		}
	}
	return 0, nil
}

func test1() {
	fmt.Println("Test 1")
	v1 := "10.20"
	v2 := "10.13"

	result, err := CompareVersion(v1, v2)
	if err != nil {
		fmt.Printf("Error %v\n", err.Error())
		return
	}

	if result == -1 {
		fmt.Println("Version 1 is latest")
	} else if result == 1 {
		fmt.Println("Version 2 is latest")
	} else {
		fmt.Println("Version 1 and v2 are same")
	}
}

func test2() {
	fmt.Println("Test 2")
	v1 := "10.20a"
	v2 := "10.13"

	result, err := CompareVersion(v1, v2)
	if err != nil {
		fmt.Printf("Error %v\n", err.Error())
		return
	}

	if result == -1 {
		fmt.Println("Version 1 is latest")
	} else if result == 1 {
		fmt.Println("Version 2 is latest")
	} else {
		fmt.Println("Version 1 and v2 are same")
	}
}

func test3() {
	fmt.Println("Test 3")
	v1 := "10.20"
	v2 := "10.13a"

	result, err := CompareVersion(v1, v2)
	if err != nil {
		fmt.Printf("Error %v\n", err.Error())
		return
	}

	if result == -1 {
		fmt.Println("Version 1 is latest")
	} else if result == 1 {
		fmt.Println("Version 2 is latest")
	} else {
		fmt.Println("Version 1 and v2 are same")
	}
}

func test4() {
	fmt.Println("Test 4")
	v1 := "10.20.11"
	v2 := "10.13"

	result, err := CompareVersion(v1, v2)
	if err != nil {
		fmt.Printf("Error %v\n", err.Error())
		return
	}

	if result == -1 {
		fmt.Println("Version 1 is latest")
	} else if result == 1 {
		fmt.Println("Version 2 is latest")
	} else {
		fmt.Println("Version 1 and v2 are same")
	}
}

func test5() {
	fmt.Println("Test 5")
	v1 := "10.20.11"
	v2 := "11.13"

	result, err := CompareVersion(v1, v2)
	if err != nil {
		fmt.Printf("Error %v\n", err.Error())
		return
	}

	if result == -1 {
		fmt.Println("Version 1 is latest")
	} else if result == 1 {
		fmt.Println("Version 2 is latest")
	} else {
		fmt.Println("Version 1 and v2 are same")
	}
}

func test6() {
	fmt.Println("Test 6")
	v1 := "10.10"
	v2 := "10.10.11"

	result, err := CompareVersion(v1, v2)
	if err != nil {
		fmt.Printf("Error %v\n", err.Error())
		return
	}

	if result == -1 {
		fmt.Println("Version 1 is latest")
	} else if result == 1 {
		fmt.Println("Version 2 is latest")
	} else {
		fmt.Println("Version 1 and v2 are same")
	}
}

func main() {
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
}
