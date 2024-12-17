package main

import (
	"fmt"
)

func printAllSubstrings(input string) {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j <= len(input); j++ { // Start j from i+1 and keep j <= len(input)
			substr := input[i:j]
			fmt.Println(substr)
		}
	}
}

func main() {
	printAllSubstrings("abcde")
}
