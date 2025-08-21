package validBraces

import "fmt"

func ValidBraces(str string) bool {
	stack := []rune{}
	runes := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, r := range str {
		fmt.Println(stack)
		switch r {
		case '(', '[', '{':
			stack = append(stack, r)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != runes[r] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
