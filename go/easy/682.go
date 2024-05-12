package easy

import "strconv"

func calPoints(operations []string) int {
	scores := make([]int, 0)
	for _, op := range operations {
		switch op {
		case "+":
			scores = append(scores, scores[len(scores)-1]+scores[len(scores)-2])
		case "D":
			scores = append(scores, 2*scores[len(scores)-1])
		case "C":
			scores = scores[:len(scores)-1]
		default:
			score, _ := strconv.Atoi(op)
			scores = append(scores, score)
		}
	}

	total := 0
	for _, s := range scores {
		total += s
	}
	return total
}
