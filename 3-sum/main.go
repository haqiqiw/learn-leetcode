package main

import (
	"fmt"
	"slices"
)

func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	outputs := [][]int{}

	for i, n := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		if nums[i] > 0 {
			break
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := n + nums[left] + nums[right]

			if sum == 0 {
				outputs = append(outputs, []int{n, nums[left], nums[right]})
				// break

				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			}
		}
	}

	return outputs
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 1, 1}))
	fmt.Println(threeSum([]int{0, 0, 0}))
	fmt.Println(threeSum([]int{0, 0, 0, 0}))
}
