package meeting

import (
	"fmt"
	"strings"
)

func UpperCase(text string) string {
	return strings.ToUpper(text)
}

func Formating(text []string) []string {
	var s []string

	for _, people := range text {
		name := strings.Split(people, ":")
		format := fmt.Sprintf("(%s, %s)", name[1], name[0])
		s = append(s, format)
	}
	return s
}
