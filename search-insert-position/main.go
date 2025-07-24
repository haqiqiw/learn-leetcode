package main

import "fmt"

func searchInsert(nums []int, target int) int {
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

	return l
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
}
