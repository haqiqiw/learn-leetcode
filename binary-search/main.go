package main

import "fmt"

func search(nums []int, target int) int {
	n := len(nums)
	l := 0
	r := n - 1

	for l <= r {
		m := l + ((r - l) / 2)

		if nums[m] == target {
			return m
		} else if target < nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return -1
}

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))
}
