package main

import "fmt"

func maxArea(height []int) int {
	n := len(height)

	left := 0
	right := n - 1

	maxPl := 0

	for left <= right {
		pl := min(height[left], height[right]) * (right - left)

		if height[left] < height[right] {
			left++
		} else if height[left] > height[right] {
			right--
		} else {
			left++
			right--
		}

		if pl > maxPl {
			maxPl = pl
		}
	}

	return maxPl
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
}
