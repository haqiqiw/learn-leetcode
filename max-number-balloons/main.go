package main

import "fmt"

// balloons
// b = 1
// a = 1
// ll = 2
// oo = 2
// n = 1
func maxNumberOfBalloons(text string) int {
	cMap := make(map[rune]int)

	for _, t := range text {
		cMap[t]++
	}

	max := 10001

	max = min(max, cMap['b']/1)
	max = min(max, cMap['a']/1)
	max = min(max, cMap['l']/2)
	max = min(max, cMap['o']/2)
	max = min(max, cMap['n']/1)

	return max
}

func main() {
	fmt.Println(maxNumberOfBalloons("nlaebolko"))
	fmt.Println(maxNumberOfBalloons("loonbalxballpoon"))
	fmt.Println(maxNumberOfBalloons("leetcode"))
	fmt.Println(maxNumberOfBalloons("abccba"))
}
