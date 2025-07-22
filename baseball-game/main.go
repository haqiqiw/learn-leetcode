package main

import (
	"fmt"
	"strconv"
)

func calPoints(operations []string) int {
	records := []int{}
	score := 0

	for _, o := range operations {
		os := string(o)
		num, err := strconv.Atoi(os)
		if err != nil {
			switch os {
			case "+":
				n := len(records)
				if n >= 2 {
					lastIdx := n - 1
					sum := records[lastIdx-1] + records[lastIdx]
					records = append(records, sum)
					score += sum
				}
			case "D":
				n := len(records)
				if n >= 1 {
					lastIdx := n - 1
					mult := records[lastIdx] * 2
					records = append(records, mult)
					score += mult
				}
			case "C":
				n := len(records)
				if n >= 1 {
					lastIdx := n - 1
					lastRec := records[lastIdx]
					records = records[:n-1]
					score -= lastRec
				}
			}
		} else {
			records = append(records, num)
			score += num
		}
	}

	return score
}

func main() {
	fmt.Println(calPoints([]string{"5", "2", "C", "D", "+"}))
	fmt.Println(calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"}))
	fmt.Println(calPoints([]string{"1", "C"}))
}
