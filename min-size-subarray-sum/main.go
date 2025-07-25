package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	l := 0
	sum := 0
	minLength := len(nums) + 1

	for r, n := range nums {
		sum += n

		for sum >= target {
			minLength = min(minLength, (r-l)+1)

			sum -= nums[l]
			l++
		}
	}

	if minLength == len(nums)+1 {
		return 0
	}

	return minLength
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}
