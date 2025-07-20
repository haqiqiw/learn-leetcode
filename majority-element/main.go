package main

import "fmt"

func majorityElement(nums []int) int {
	num, count := 0, 0

	for _, n := range nums {
		if count == 0 {
			num = n
		}

		if n == num {
			count++
		} else {
			count--
		}
	}

	count = 0
	for _, n := range nums {
		if n == num {
			count++
		}
	}

	if count > len(nums)/2 {
		return num
	}

	return -1
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
