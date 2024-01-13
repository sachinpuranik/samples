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

// Algo
// Run the program and observe the output.
// I represents the wor number.
//

func convertZigZagFast(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var result strings.Builder
	n := len(s)
	cycleLen := 2 * (numRows - 1)

	for i := 0; i < numRows; i++ {
		for j := 0; j+i < n; j += cycleLen {
			fmt.Printf("\n result.WriteByte(s[j(%d)+i(%d)])) : %d", j, i, j+i)
			result.WriteByte(s[j+i])
			if i != 0 && i != numRows-1 && j+cycleLen-i < n {
				fmt.Printf("\n result.WriteByte(s[j(%d)+cycleLen(%d)-i(%d)]) : %d", j, cycleLen, i, j+cycleLen-i)
				result.WriteByte(s[j+cycleLen-i])
			}
		}
	}

	return result.String()
}

func main() {
	s := "PAYPALISHIRING"
	numRows := 3

	out := convertZigZagSlow(s, numRows)

	if out == "PAHNAPLSIIGYIR" {
		fmt.Println("pass")
	} else {
		fmt.Println("fail")
	}

	numRows++
	out = convertZigZagFast(s, numRows)

	if out == "PINALSIGYAHRPI" {
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
