package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	l := 0
	r := (m * n) - 1

	for l <= r {
		m := l + ((r - l) / 2)

		row := m / n
		col := m % n
		v := matrix[row][col]

		if v == target {
			return true
		} else if target < v {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return false
}

func main() {
	fmt.Println(searchMatrix(
		[][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		}, 3,
	))
	fmt.Println(searchMatrix(
		[][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		}, 13,
	))
}
