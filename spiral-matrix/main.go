package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	outputs := []int{}

	rowStart := 0
	rowEnd := len(matrix) - 1
	colStart := 0
	colEnd := len(matrix[0]) - 1

	for {
		if rowStart > rowEnd || colStart > colEnd {
			break
		}

		// 1. top left to top right
		for c := colStart; c <= colEnd; c++ {
			outputs = append(outputs, matrix[rowStart][c])
			fmt.Println("Step 1: ", outputs)
		}
		rowStart++

		if rowStart > rowEnd {
			break
		}

		// 2. top right to bottm right
		for r := rowStart; r <= rowEnd; r++ {
			outputs = append(outputs, matrix[r][colEnd])
			fmt.Println("Step 2: ", outputs)
		}
		colEnd--

		if colStart > colEnd {
			break
		}

		// 3. bottom middle right to bottom left
		for c := colEnd; c >= colStart; c-- {
			outputs = append(outputs, matrix[rowEnd][c])
			fmt.Println("Step 3: ", outputs)

		}
		rowEnd--

		if rowStart > rowEnd {
			break
		}

		// 4. bottom left to bottom top
		for r := rowEnd; r >= rowStart; r-- {
			outputs = append(outputs, matrix[r][colStart])
			fmt.Println("Step 4: ", outputs)
		}
		colStart++
	}

	return outputs
}

func main() {
	fmt.Println(spiralOrder(
		[][]int{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{7, 8, 9},
		},
	))
	fmt.Println(spiralOrder(
		[][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		},
	))
}
