package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	stacks := []int{}
	outputs := make([]int, len(temperatures))

	for i, t := range temperatures {
		for len(stacks) > 0 && temperatures[stacks[len(stacks)-1]] < t {
			prevIdx := stacks[len(stacks)-1]
			stacks = stacks[:len(stacks)-1]
			outputs[prevIdx] = i - prevIdx
		}

		stacks = append(stacks, i)
	}

	return outputs
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
