package main

import "fmt"

func isSubsequence(s string, t string) bool {
	si := 0

	if len(s) == 0 {
		return true
	}

	for i := 0; i < len(t); i++ {
		v := t[i]

		if v == s[si] {
			si++
		}

		if si == (len(s)) {
			return true
		}
	}

	return si == (len(s))
}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
}
