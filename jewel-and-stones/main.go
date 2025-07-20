package main

import "fmt"

func numJewelsInStones(jewels string, stones string) int {
	jMap := make(map[rune]struct{})
	for _, j := range jewels {
		jMap[j] = struct{}{}
	}

	count := 0

	for _, s := range stones {
		if _, exist := jMap[s]; exist {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
	fmt.Println(numJewelsInStones("z", "ZZ"))
}
