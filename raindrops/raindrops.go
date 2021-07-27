package raindrops

import "strconv"

func Convert(n int) string {
	converted := ""

	if n%3 == 0 {
		converted += "Pling"
	}
	if n%5 == 0 {
		converted += "Plang"
	}
	if n%7 == 0 {
		converted += "Plong"
	}
	if converted == "" {
		converted = strconv.Itoa(n)
	}

	return converted
}
