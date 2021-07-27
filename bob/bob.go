// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!

	isWord := false
	for _, c := range remark {
		if unicode.IsLetter(c) {
			isWord = true
		}
	}
	trimedRemark := strings.TrimSpace(remark)
	remarkUpper := strings.ToUpper(trimedRemark)
	isQuestion := strings.HasSuffix(trimedRemark, "?")

	if remarkUpper == remark && isQuestion && isWord {
		return "Calm down, I know what I'm doing!"
	}
	if remarkUpper == remark && isWord {
		return "Whoa, chill out!"
	}
	if isQuestion {
		return "Sure."
	}
	if !isWord && trimedRemark == "" {
		return "Fine. Be that way!"
	}

	return "Whatever."
}
