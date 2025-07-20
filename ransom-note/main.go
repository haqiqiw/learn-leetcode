package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	nMap := make(map[rune]int)
	mMap := make(map[rune]int)

	for _, n := range ransomNote {
		nMap[n]++
	}

	for _, m := range magazine {
		mMap[m]++
	}

	for n, t := range nMap {
		mt, exist := mMap[n]
		if !exist || mt < t {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
	fmt.Println(canConstruct("a", "aab"))
}
