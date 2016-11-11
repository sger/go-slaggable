package sluggable

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

const LOWERCASE_NUMBERS_DASHES = "([^A-Za-z0-9]|-)+"
const SEPARATOR = "-"

func Slugify(s string) string {
	reg, err := regexp.Compile(LOWERCASE_NUMBERS_DASHES)
	if err != nil {
		log.Fatal(err)
	}

	safe := reg.ReplaceAllString(s, "-")
	safe = strings.ToLower(strings.Trim(safe, SEPARATOR))
	fmt.Println("result ", safe)
	return safe
}
