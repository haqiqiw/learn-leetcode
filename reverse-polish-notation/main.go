package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stacks := []int{}

	for _, t := range tokens {
		num, err := strconv.Atoi(t)
		if err != nil {
			n := len(stacks)
			lastIdx := n - 1

			op := string(t)
			switch op {
			case "+":
				res := stacks[lastIdx-1] + stacks[lastIdx]
				stacks = stacks[:n-2]
				stacks = append(stacks, res)
			case "-":
				res := stacks[lastIdx-1] - stacks[lastIdx]
				stacks = stacks[:n-2]
				stacks = append(stacks, res)
			case "*":
				res := stacks[lastIdx-1] * stacks[lastIdx]
				stacks = stacks[:n-2]
				stacks = append(stacks, res)
			case "/":
				res := stacks[lastIdx-1] / stacks[lastIdx]
				stacks = stacks[:n-2]
				stacks = append(stacks, res)
			default:
				continue
			}
		} else {
			stacks = append(stacks, num)
		}
	}

	return stacks[0]
}

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
