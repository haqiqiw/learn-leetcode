package main

import "fmt"

func search(nums []int, target int) int {
	n := len(nums)
	l := 0
	r := n - 1

	for l <= r {
		m := l + ((r - l) / 2) // 3[7]

		if nums[m] == target {
			return m
		} else if nums[l] <= nums[m] { // left to mid sorted
			if target >= nums[l] && target < nums[m] {
				// target on left side
				r = m - 1
			} else {
				// target on right side
				l = m + 1
			}
		} else if nums[l] > nums[m] { // mid to right sorted
			if target > nums[m] && target <= nums[r] {
				// target on right side
				l = m + 1
			} else {
				// target on left side
				r = m - 1
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(search([]int{6, 7, 0, 1, 2, 3, 4, 5}, 3))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	fmt.Println(search([]int{1}, 0))
}
