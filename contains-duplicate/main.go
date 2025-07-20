package main

import "fmt"

func containsDuplicate(nums []int) bool {
	nMap := make(map[int]struct{})

	for _, n := range nums {
		if _, exist := nMap[n]; exist {
			return true
		} else {
			nMap[n] = struct{}{}
		}
	}

	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
