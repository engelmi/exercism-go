package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	seenRunesMap := make(map[rune]bool)
	for _, c := range strings.ToLower(word) {
		_, found := seenRunesMap[c]
		if found && unicode.IsLetter(c) {
			return false
		}
		seenRunesMap[c] = true
	}
	return true
}
