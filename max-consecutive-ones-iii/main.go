package main

import "fmt"

func longestOnes(nums []int, k int) int {
	n := len(nums)
	l := 0

	maxLength := 0
	zero := 0

	for r := 0; r < n; r++ {
		if nums[r] == 0 {
			zero++
		}

		for zero > k {
			if nums[l] == 0 {
				zero--
			}
			l++
		}

		length := r - l + 1
		maxLength = max(maxLength, length)
	}

	return maxLength
}

func main() {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
}
