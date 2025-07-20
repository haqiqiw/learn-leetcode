package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := make(map[rune]int)

	for _, v := range s {
		sMap[v]++
	}

	for _, v := range t {
		vt, exist := sMap[v]
		if exist {
			if vt <= 0 {
				return false
			} else {
				sMap[v]--
			}
		} else {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}
