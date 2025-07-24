package main

import "fmt"

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}

	l := 1
	r := num / 2

	for l <= r {
		m := l + ((r - l) / 2)
		sq := m * m

		if sq == num {
			return true
		} else if sq < num {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return false
}

func main() {
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquare(14))
	fmt.Println(isPerfectSquare(1))
}
