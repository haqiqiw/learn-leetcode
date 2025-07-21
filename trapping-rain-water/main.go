package main

import "fmt"

func trap(height []int) int {
	n := len(height)
	output := 0

	maxLefts := make([]int, n)
	maxRights := make([]int, n)

	for i := 0; i <= n-1; i++ {
		if i == 0 {
			maxLefts[i] = height[i]
		} else {
			maxLefts[i] = max(maxLefts[i-1], height[i])
		}
	}

	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			maxRights[i] = height[i]
		} else {
			maxRights[i] = max(maxRights[i+1], height[i])
		}
	}

	for i := 0; i < n-1; i++ {
		wall := min(maxLefts[i], maxRights[i])
		trap := wall - height[i]
		if trap > 0 {
			output += trap
		}
	}

	return output
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
