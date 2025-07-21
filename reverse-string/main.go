package main

import "fmt"

func reverseString(s []byte) {
	if len(s) == 0 || len(s) == 1 {
		return
	}

	left := 0
	right := len(s) - 1
	for {
		fmt.Println(left, right)
		if left >= right {
			break
		}

		temp := s[left]
		s[left] = s[right]
		s[right] = temp

		left++
		right--
	}
}

func main() {
	s1 := []string{"h", "e", "l", "l", "o"}
	b1 := []byte{}
	for _, s := range s1 {
		b1 = append(b1, s[0])
	}
	fmt.Println(b1)
	reverseString(b1)
	fmt.Println(b1)
}
