package scrabble

import (
	"strings"
)

var valueMap = map[string]int{
	"AEIOULNRST": 1,
	"DG":         2,
	"BCMP":       3,
	"FHVWY":      4,
	"K":          5,
	"JX":         8,
	"QZ":         10,
}

func Score(word string) int {
	score := 0
	for _, c := range strings.ToUpper(word) {
		for k, v := range valueMap {
			if strings.Contains(k, string(c)) {
				score += v
			}
		}
	}

	return score
}
