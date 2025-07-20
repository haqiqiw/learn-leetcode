package main

import (
	"fmt"
	"slices"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{{strs[0]}}
	}

	gMap := make(map[string][]string)

	for _, s := range strs {
		ssort := strings.Split(s, "")
		slices.Sort(ssort)

		sjoin := strings.Join(ssort, "")

		_, exist := gMap[sjoin]
		if exist {
			gMap[sjoin] = append(gMap[sjoin], s)
		} else {
			gMap[sjoin] = []string{s}
		}
	}

	outputs := [][]string{}
	for _, v := range gMap {
		outputs = append(outputs, v)
	}

	return outputs
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{""}))
	fmt.Println(groupAnagrams([]string{"a"}))
}
