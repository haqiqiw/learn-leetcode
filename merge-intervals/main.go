package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	n := len(intervals)

	if n == 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	start := intervals[0][0]
	end := intervals[0][1]

	outputs := [][]int{}

	for i := 0; i < n; i++ {
		hasNext := (i + 1) < n

		if hasNext {
			if intervals[i+1][0] <= end {
				end = max(end, intervals[i+1][1])
				continue
			} else {
				outputs = append(outputs, []int{start, end})
				start = intervals[i+1][0]
				end = intervals[i+1][1]
			}
		} else {
			outputs = append(outputs, []int{start, end})
		}
	}

	return outputs
}

func main() {
	fmt.Println(merge([][]int{
		[]int{2, 6},
		[]int{1, 3},
		[]int{15, 18},
		[]int{8, 10},
	}))
	fmt.Println(merge([][]int{
		[]int{4, 5},
		[]int{1, 4},
	}))
	fmt.Println(merge([][]int{
		[]int{1, 10},
		[]int{2, 3},
		[]int{4, 5},
	}))
}
