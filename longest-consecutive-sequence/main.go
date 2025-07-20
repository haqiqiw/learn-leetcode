package main

import "fmt"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	nMap := make(map[int]bool)
	max := 0

	for _, n := range nums {
		nMap[n] = true
	}

	for _, n := range nums {
		// check prev
		if !nMap[n-1] {
			// starting point
			curNum := n
			iterMax := 1

			// check next
			for nMap[curNum+1] {
				curNum++
				iterMax++
			}

			if iterMax > max {
				max = iterMax
			}
		}
	}

	return max
}

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Println(longestConsecutive([]int{1, 0, 1, 2}))
}
