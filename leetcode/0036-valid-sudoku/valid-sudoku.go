package leetcode

import (
	"bytes"
	"fmt"
)

func isValidSudoku(board [][]byte) bool {

	cols := make(map[int][]byte)
	rows := make(map[int][]byte)
	boxes := make(map[string][]byte)

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			cur := board[row][col]
			if cur != '.' {
				if bytes.Contains(rows[row], []byte{cur}) {
					return false
				}

				if bytes.Contains(cols[col], []byte{cur}) {
					return false
				}

				if bytes.Contains(boxes[boxKey(row, col)], []byte{cur}) {
					return false
				}

				rows[row] = append(rows[row], cur)
				cols[col] = append(cols[col], cur)
				boxes[boxKey(row, col)] = append(boxes[boxKey(row, col)], cur)
			}
		}
	}

	return true
}

func boxKey(row, col int) string {
	return (fmt.Sprint(row/3) + " " + fmt.Sprint(col/3))
}
