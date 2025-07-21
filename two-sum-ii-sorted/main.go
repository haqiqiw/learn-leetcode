package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left <= right {
		sum := numbers[left] + numbers[right]

		if sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			return []int{left + 1, right + 1}
		}
	}

	return nil
}

func main() {
	fmt.Println(twoSum([]int{5, 25, 75}, 100))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 3, 4}, 6))
	fmt.Println(twoSum([]int{-1, 0}, -1))
	fmt.Println(twoSum([]int{10, 20, 30, 40, 50}, 70))
}
