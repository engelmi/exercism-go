package luhn

import (
	"strconv"
	"strings"
)

func Valid(idStr string) bool {
	trimmedIdStr := strings.ReplaceAll(idStr, " ", "")
	if len(trimmedIdStr) < 2 {
		return false
	}

	isSecond := false
	sum := 0
	for i := len(trimmedIdStr) - 1; i >= 0; i-- {
		v, err := strconv.Atoi(string(trimmedIdStr[i]))
		if err != nil {
			return false
		}

		if isSecond {
			vd := v * 2
			if vd > 9 {
				vd = vd - 9
			}
			sum += vd
			isSecond = false
		} else {
			sum += v
			isSecond = true
		}
	}
	return sum%10 == 0
}
