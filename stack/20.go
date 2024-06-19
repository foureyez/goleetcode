package stack

import "fmt"

func isValid(s string) bool {
	stack := make([]rune, 0)

	for _, p := range s {
		if p == '(' || p == '{' || p == '[' {
			stack = append(stack, p)
		} else if len(stack) > 0 {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if !isEquivalentParenthesis(pop, p) {
				return false
			}
		} else {
			return false
		}
	}
	return len(stack) == 0
}

func isEquivalentParenthesis(left, right rune) bool {
	pair := fmt.Sprintf("%c%c", left, right)
	if pair == "()" || pair == "{}" || pair == "[]" {
		return true
	}
	return false
}
