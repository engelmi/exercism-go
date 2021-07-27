package twofer

import "fmt"

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	target := "you"
	if name != "" {
		target = name
	}

	return fmt.Sprintf("One for %s, one for me.", target)
}
