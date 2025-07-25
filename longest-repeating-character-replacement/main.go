package main

import "fmt"

func characterReplacement(s string, k int) int {
	l := 0
	maxFreq := 0
	maxLength := 0
	cMap := map[rune]int{}

	for r, c := range s {
		cMap[c]++
		maxFreq = max(maxFreq, cMap[c])

		for ((r-l)+1)-maxFreq > k {
			cMap[rune(s[l])]--
			l++
		}

		length := (r - l) + 1
		maxLength = max(maxLength, length)
	}

	return maxLength
}

func main() {
	fmt.Println(characterReplacement("ABAB", 2))
	fmt.Println(characterReplacement("AABABBA", 1))
}
