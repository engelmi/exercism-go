package robotname

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	rand.Seed(time.Now().UnixNano())

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	nums := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	charLength := 2
	numLength := 3
	var b strings.Builder
	for i := 0; i < charLength; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	for i := 0; i < numLength; i++ {
		b.WriteString(strconv.Itoa(nums[rand.Intn(len(nums))]))
	}

	r.name = b.String()
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}
