package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	outputs := make([]int, len(nums))

	curLeft := 1
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			outputs[i] = curLeft
		} else {
			outputs[i] = curLeft * nums[i-1]
			curLeft = outputs[i]
		}
	}

	curRight := 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			outputs[i] *= curRight
		} else {
			temp := curRight * nums[i+1]
			curRight = temp
			outputs[i] *= temp
		}
	}

	return outputs
}

func main() {
	// fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	// fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
	fmt.Println(productExceptSelf([]int{5, 2, 3, 4}))
}
