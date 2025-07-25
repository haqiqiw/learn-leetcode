package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Map := map[rune]int{}
	s2Map := map[rune]int{}

	for _, c := range s1 {
		s1Map[c]++
	}

	for i := 0; i < len(s1); i++ {
		s2Map[rune(s2[i])]++
	}

	if checkMap(s1Map, s2Map) {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		s2Map[rune(s2[i])]++
		s2Map[rune(s2[i-len(s1)])]--

		if checkMap(s1Map, s2Map) {
			return true
		}
	}

	return false
}

func checkMap(s1Map map[rune]int, s2Map map[rune]int) bool {
	for k, v := range s1Map {
		if v != s2Map[k] {
			return false
		}
	}

	return true
}

func checkInclusionSlow(s1 string, s2 string) bool {
	s1Map := map[rune]int{}
	s1Len := len(s1)
	s2len := len(s2)

	for _, c := range s1 {
		s1Map[c]++
	}

	l := 0
	r := s1Len - 1

	for r < s2len {
		s2Map := map[rune]int{}
		for _, c := range s2[l : r+1] {
			s2Map[c]++
		}

		valid := true
		for k, v := range s2Map {
			if s1Map[k] != v {
				valid = false
			}
		}

		if valid {
			return true
		}

		l++
		r++
	}

	return false
}

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
	fmt.Println(checkInclusion("ab", "eidboaoo"))
}
