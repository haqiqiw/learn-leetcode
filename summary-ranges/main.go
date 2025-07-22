package main

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	outputs := []string{}
	initRange := nums[0]

	for i, v := range nums {
		hasNext := (i + 1) < len(nums)

		if hasNext && nums[i+1] == v+1 {
			continue
		} else {
			if initRange == v {
				outputs = append(outputs, fmt.Sprint(v))
			} else {
				outputs = append(outputs, fmt.Sprintf("%d->%d", initRange, v))
			}

			if hasNext {
				initRange = nums[i+1]
			}
		}
	}

	return outputs
}

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}
