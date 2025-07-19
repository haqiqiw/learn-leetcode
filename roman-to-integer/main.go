package main

import (
	"fmt"
	"strings"
)

var romanMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	r := strings.Split(s, "")
	sum := 0

	skipNext := false

	for i, v := range r {
		if skipNext {
			skipNext = false
			continue
		}

		hasNext := (i + 1) < len(r)

		switch v {
		case "C":
			if hasNext {
				switch r[i+1] {
				case "M":
					sum += 900
					skipNext = true
				case "D":
					sum += 400
					skipNext = true
				default:
					sum += romanMap[v]
				}
			} else {
				sum += romanMap[v]
			}
		case "X":
			if hasNext {
				switch r[i+1] {
				case "C":
					sum += 90
					skipNext = true
				case "L":
					sum += 40
					skipNext = true
				default:
					sum += romanMap[v]
				}
			} else {
				sum += romanMap[v]
			}
		case "I":
			if hasNext {
				switch r[i+1] {
				case "X":
					sum += 9
					skipNext = true
				case "V":
					sum += 4
					skipNext = true
				default:
					sum += romanMap[v]
				}
			} else {
				sum += romanMap[v]
			}
		default:
			sum += romanMap[v]
		}
	}

	return sum
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
	fmt.Println(romanToInt("MCDLXXVI"))
}
