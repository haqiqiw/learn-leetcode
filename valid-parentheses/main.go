package main

import "fmt"

func isValid(s string) bool {
	stacks := []string{}

	for _, v := range s {
		n := len(stacks)
		lastIdx := n - 1
		sv := string(v)
		switch sv {
		case "(":
			stacks = append(stacks, sv)
		case "{":
			stacks = append(stacks, sv)
		case "[":
			stacks = append(stacks, sv)
		case ")":
			if n >= 1 {
				if stacks[lastIdx] == "(" {
					stacks = stacks[:n-1]
				} else {
					return false
				}
			} else {
				return false
			}
		case "}":
			if n >= 1 {
				if stacks[lastIdx] == "{" {
					stacks = stacks[:n-1]
				} else {
					return false
				}
			} else {
				return false
			}
		case "]":
			if n >= 1 {
				if stacks[lastIdx] == "[" {
					stacks = stacks[:n-1]
				} else {
					return false
				}
			} else {
				return false
			}
		}
	}

	return len(stacks) == 0
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([])"))
}
