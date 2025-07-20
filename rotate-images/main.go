package main

import "fmt"

func rotate(matrix [][]int) {
	print(matrix)
	rows := len(matrix) - 1

	// transpose
	for i := 0; i <= rows; i++ {
		for j := i + 1; j <= rows; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}
	print(matrix)

	// reverse
	for i := 0; i <= rows; i++ {
		reverse(matrix[i])
	}
	print(matrix)
}

func reverse(row []int) {
	start := 0
	end := len(row) - 1

	for {
		if start >= end {
			break
		}

		temp := row[start]
		row[start] = row[end]
		row[end] = temp

		start++
		end--
	}
}

func main() {
	rotate([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
	rotate([][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	})
}

func print(matrix [][]int) {
	if len(matrix) == 0 {
		fmt.Println("[]")
		return
	}

	for r := 0; r < len(matrix); r++ {
		fmt.Println(matrix[r])
	}
}
