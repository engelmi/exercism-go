package acronym

import (
	"log"
	"regexp"
	"strings"
)

func Abbreviate(s string) string {
	abbr := ""

	reg, err := regexp.Compile("([a-zA-Z0-9]+)('[a-z])?")
	if err != nil {
		log.Fatal(err)
	}

	for _, word := range reg.FindAllString(s, 100) {
		abbr += strings.ToUpper(word[0:1])
	}

	return abbr
}
