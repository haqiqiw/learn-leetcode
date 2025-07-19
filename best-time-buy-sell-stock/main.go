package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for _, p := range prices {
		curProfit := p - minPrice
		maxProfit = max(maxProfit, curProfit)
		minPrice = min(minPrice, p)
	}

	return maxProfit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
