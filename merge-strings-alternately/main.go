package main

import (
	"fmt"
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {
	i := 0

	n1 := len(word1)
	n2 := len(word2)

	s1 := strings.Split(word1, "")
	s2 := strings.Split(word2, "")

	var sb strings.Builder

	for {
		if i < n1 {
			sb.WriteString(s1[i])
		}

		if i < n2 {
			sb.WriteString(s2[i])
		}

		if i >= n1 && i >= n2 {
			break
		}

		i++
	}

	return sb.String()
}

func main() {
	fmt.Println(mergeAlternately("abc", "pqr"))
	fmt.Println(mergeAlternately("ab", "pqrs"))
	fmt.Println(mergeAlternately("abcd", "pq"))
}
