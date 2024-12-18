package main

// ref : https://leetcode.com/list/du693s/
import (
	"fmt"
	"strings"
)

func convertZigZagSlow(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([][]string, numRows)

	rising := true
	risingCount := 0
	for _, r := range s {
		rows[risingCount] = append(rows[risingCount], string(r))

		if rising {
			if risingCount == numRows-1 {
				risingCount--
				rising = false
				continue
			}
			risingCount++
			continue
		}
		if risingCount == 0 {
			rising = true
			risingCount++
			continue
		}
		risingCount--
	}
	final := ""
	for _, row := range rows {
		for _, item := range row {
			final = final + item
		}
	}

	return final
}

// Given string "ABCDEFGHIJKLMN" and lets say number of rows is 3. In this case, the string will be written in the following way:
// A   E   I   M
// B D F H J L N
// C   G   K
// Same string will be written in the following way if number of rows is 4:
// A     G     M
// B   F H   L N
// C E   I K
// D     J

// Algo
// 1. Create a string builder to store the result.
// 2. Find the cycle length. It is 2 * (numRows - 1). Essentially span how many characters are there in a cycle.
// 3. For each row , iterate through array pick the elements. if row number is 0 or row number numRows-1, then we will pick only one element j+i , else we will pick two elements j+i and j+cycleLen-i (its like v)
// 4. example - ABCDEFGHIJKLMN with 3 rows, for row 0 , it will be A, E, I, M., for row 1, it will be B, D, F, H, J, L, N. for row 2, it will be C, G, K.
// Note - In each cycle

func convertZigZag(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	var result strings.Builder
	n := len(s)
	cycleLen := 2 * (numRows - 1) // Full cycle length

	for row := 0; row < numRows; row++ {
		for j := row; j < n; j += cycleLen {
			// Add the primary character in the current row
			result.WriteByte(s[j])

			// Add the diagonal character for rows that are not the first or last
			diag := j + cycleLen - 2*row
			if row > 0 && row < numRows-1 && diag < n {
				result.WriteByte(s[diag])
			}
		}
	}

	return result.String()
}

func main() {
	s := "ABCDEFGHIJKLMN"

	numRows := 3

	out := convertZigZagSlow(s, numRows)

	if out == "AEIMBDFHJLNCGK" {
		fmt.Println("pass")
	} else {
		fmt.Println("fail")
	}

	numRows = 4
	out = convertZigZag(s, numRows)

	if out == "AGMBFHLNCEIKDJ" {
		fmt.Println("pass")
	} else {
		fmt.Println("fail")
	}

	// out = convertZigZagFast("AB", 1)

	// if out == "AB" {
	// 	fmt.Println("pass")
	// } else {
	// 	fmt.Println("fail")
	// }
}
