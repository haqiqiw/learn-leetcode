package main

import (
	"fmt"
)

func isValidSudoku(board [][]byte) bool {
	rows := make(map[int]map[string]struct{})
	cols := make(map[int]map[string]struct{})
	boxs := make(map[int]map[string]struct{})

	for i := 0; i <= len(board)-1; i++ {
		for j := 0; j <= len(board[0])-1; j++ {
			num := string(board[i][j])
			if num == "." {
				continue
			}

			// row
			if _, rexist := rows[i][num]; rexist {
				return false
			} else {
				if rows[i] == nil {
					rows[i] = map[string]struct{}{}
				}
				rows[i][num] = struct{}{}
			}

			// col
			if _, cexist := cols[j][num]; cexist {
				return false
			} else {
				if cols[j] == nil {
					cols[j] = map[string]struct{}{}
				}
				cols[j][num] = struct{}{}
			}

			// box
			bi := ((i / 3) * 3) + (j / 3)
			if _, bexist := boxs[bi][num]; bexist {
				return false
			} else {
				if boxs[bi] == nil {
					boxs[bi] = map[string]struct{}{}
				}
				boxs[bi][num] = struct{}{}
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isValidSudoku(build([][]string{
		{"5", "3", ".", ".", "7", ".", ".", ".", "."},
		{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", "6", "."},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	})))
}

func build(board [][]string) [][]byte {
	rows := len(board)
	cols := len(board[0])
	boardAsByte := make([][]byte, rows)
	for i := range boardAsByte {
		boardAsByte[i] = make([]byte, cols)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			boardAsByte[r][c] = board[r][c][0]
		}
	}

	return boardAsByte
}
