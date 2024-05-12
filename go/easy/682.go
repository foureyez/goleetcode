package easy

import "strconv"

func calPoints(operations []string) int {
	stack := make([]int, 0)
	for _, op := range operations {
		switch op {
		case "+":
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
		case "D":
			stack = append(stack, 2*stack[len(stack)-1])
		case "C":
			stack = stack[:len(stack)-1]
		default:
			score, _ := strconv.Atoi(op)
			stack = append(stack, score)
		}
	}

	total := 0
	for _, s := range stack {
		total += s
	}
	return total
}
