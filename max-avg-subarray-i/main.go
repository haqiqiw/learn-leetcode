package main

import "fmt"

func findMaxAverage(nums []int, k int) float64 {
	n := len(nums)
	curSum := float64(0)
	curAvg := float64(0)

	for i := 0; i < k; i++ {
		curSum += float64(nums[i])
	}

	curAvg = float64(curSum) / float64(k)

	for i := k; i < n; i++ {
		curSum += float64(nums[i])
		curSum -= float64(nums[i-k])

		avg := float64(curSum) / float64(k)
		curAvg = max(curAvg, avg)
	}

	return curAvg
}

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
	fmt.Println(findMaxAverage([]int{5}, 1))
}
