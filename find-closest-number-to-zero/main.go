package main

import (
	"fmt"
	"math"
)

func findClosestNumber(nums []int) int {
	curNum := nums[0]
	curMax := int(math.Abs(float64(nums[0])))
	for _, v := range nums {
		d := int(math.Abs(float64(v)))

		if d < curMax {
			curNum = v
			curMax = d
		} else if d == curMax {
			if v > curNum {
				curNum = v
			}
		}
	}
	return curNum
}

func main() {
	fmt.Println(findClosestNumber([]int{-4, -2, 1, 4, 8}))
	fmt.Println(findClosestNumber([]int{2, -1, 1}))
}
