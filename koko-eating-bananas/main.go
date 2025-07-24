package main

import (
	"fmt"
	"math"
	"slices"
)

func minEatingSpeed(piles []int, h int) int {
	l := 1
	r := slices.Max(piles)
	minK := r

	for l <= r {
		m := l + ((r - l) / 2)
		sh := sumHours(piles, m)

		if sh <= h {
			minK = m
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return minK
}

func sumHours(piles []int, k int) int {
	sum := 0

	for _, p := range piles {
		sum += int(math.Ceil(float64(p) / float64(k)))
	}

	return sum
}

func main() {
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 5))
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 6))
}
