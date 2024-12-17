package main

import (
	"fmt"
)

// basic logic -
// loop through each char , or element.
//In each loop , keep the ith element aside and send the further string to get sub-combinations.
//

func generateCombinations(str string) []string {
	if len(str) == 1 {
		return []string{str}
	}

	combinations := []string{}
	for i := 0; i < len(str); i++ {
		firstChar := string(str[i])
		remaining := str[:i] + str[i+1:]
		subCombinations := generateCombinations(remaining)
		for _, subCombination := range subCombinations {
			combinations = append(combinations, firstChar+subCombination)
		}
	}
	return combinations
}

func main() {
	inputString := "abc"
	combinations := generateCombinations(inputString)
	fmt.Println(combinations) // Output: [abc acb bac bca cab cba]
}
