package main

import "fmt"

func findMin(nums []int) int {
	n := len(nums)
	l := 0
	r := n - 1

	for l < r {
		m := l + ((r - l) / 2)

		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	return nums[l]
}

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
}
