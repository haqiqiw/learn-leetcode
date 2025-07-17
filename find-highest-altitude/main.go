package main

import (
	"fmt"
)

func largestAltitude(gain []int) int {
	curAlt := 0
	maxAlt := 0

	for i := 0; i < len(gain); i++ {
		curAlt += gain[i]
		maxAlt = max(curAlt, maxAlt)
	}

	return maxAlt
}

func main() {
	fmt.Println(largestAltitude([]int{-5, 1, 5, 0, -7}))
	fmt.Println(largestAltitude([]int{-4, -3, -2, -1, 4, 3, 2}))
}
