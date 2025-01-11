// ref - https://www.youtube.com/watch?v=PMMc4VsIacU
package main

import (
	"fmt"
)

func main() {
	test := []struct {
		input    [][]byte
		expected [][]byte
	}{
		{
			input: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			},
			expected: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
		{
			input: [][]byte{
				{'O', 'X', 'O'},
				{'X', 'O', 'X'},
				{'O', 'X', 'O'},
			},
			expected: [][]byte{
				{'O', 'X', 'O'},
				{'X', 'X', 'X'},
				{'O', 'X', 'O'},
			},
		},
		{
			input: [][]byte{
				{'O', 'O', 'O', 'O', 'X', 'X'},
				{'O', 'O', 'O', 'O', 'O', 'O'},
				{'O', 'X', 'O', 'X', 'O', 'O'},
				{'O', 'X', 'O', 'O', 'X', 'O'},
				{'O', 'X', 'O', 'X', 'O', 'O'},
				{'O', 'X', 'O', 'O', 'O', 'O'},
			},

			expected: [][]byte{
				{'O', 'O', 'O', 'O', 'X', 'X'},
				{'O', 'O', 'O', 'O', 'O', 'O'},
				{'O', 'X', 'O', 'X', 'O', 'O'},
				{'O', 'X', 'O', 'O', 'X', 'O'},
				{'O', 'X', 'O', 'X', 'O', 'O'},
				{'O', 'X', 'O', 'O', 'O', 'O'},
			},
		},
	}

	// marked = make(map[xy]bool, 0)
	// visited = make(map[xy]bool, 0)

	for _, t := range test {
		matrix := t.input
		fmt.Println("input\n ", t.input)
		solve(matrix)
		fmt.Println("expected \n", t.expected)
		fmt.Println("output \n", matrix)
	}
}

func solve(board [][]byte) {
	rows := len(board)
	if rows == 0 {
		return
	}
	cols := len(board[0])

	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row < 0 || row >= rows || col < 0 || col >= cols || board[row][col] != 'O' {
			return
		}
		board[row][col] = '*'
		dfs(row+1, col)
		dfs(row-1, col)
		dfs(row, col+1)
		dfs(row, col-1)
	}

	for row := 0; row < rows; row++ {
		dfs(row, 0)
		dfs(row, cols-1)
	}
	for col := 0; col < cols; col++ {
		dfs(0, col)
		dfs(rows-1, col)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if board[row][col] == 'O' {
				board[row][col] = 'X'
			} else if board[row][col] == '*' {
				board[row][col] = 'O'
			}
		}
	}

	return
}
