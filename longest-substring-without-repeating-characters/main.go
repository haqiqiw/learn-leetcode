package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	chars := []rune(s)

	var res []rune
	var cur rune

	for i, c := range chars {
		if c == cur {
			res = []rune{c}
			continue
		}

		cur = c
	}
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
