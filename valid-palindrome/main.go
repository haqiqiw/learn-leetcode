package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	sr := []rune{}
	for _, v := range s {
		if ('a' <= v && v <= 'z') ||
			('A' <= v && v <= 'Z') ||
			('0' <= v && v <= '9') {
			sr = append(sr, []rune(strings.ToLower(string(v)))[0])
		}
	}

	if len(sr) == 0 || len(sr) == 1 {
		return true
	}
	fmt.Println(string(sr))

	left := 0
	right := len(sr) - 1

	for left <= right {
		if sr[left] != sr[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(" "))
}
