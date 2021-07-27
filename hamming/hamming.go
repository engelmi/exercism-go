package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Distance not defined for strings of unequal length")
	}

	distance := 0
	for i, _ := range a {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
