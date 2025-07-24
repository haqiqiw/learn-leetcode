package main

import "fmt"

// dummy func
func isBadVersion(v int) bool {
	return v >= 4 // tc 1
	// return v >= 1 // tc 2
	// return v >= 5 // tc myself
}

func firstBadVersion(n int) int {
	l := 1
	r := n

	for l < r {
		m := l + ((r - l) / 2)
		fmt.Println("before: ", l, m, r)

		if isBadVersion(m) {
			r = m
		} else {
			l = m + 1
		}
		fmt.Println("after: ", l, m, r)
	}

	return l
}

func main() {
	// [1, 2, 3, 4, 5]
	fmt.Println(firstBadVersion(10))
	// fmt.Println(firstBadVersion(1))
	// fmt.Println(firstBadVersion(7))
}
