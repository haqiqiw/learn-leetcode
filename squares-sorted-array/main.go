package main

import "fmt"

func sortedSquares(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nums
	}

	if n == 1 {
		return []int{nums[0] * nums[0]}
	}

	outputs := make([]int, n)
	counter := n - 1

	left := 0
	right := n - 1

	for left <= right {
		lp := nums[left] * nums[left]
		rp := nums[right] * nums[right]

		if lp > rp {
			outputs[counter] = lp
			left++
		} else {
			outputs[counter] = rp
			right--
		}
		counter--
	}

	return outputs
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))
}
