package meeting

import (
	"fmt"
	"sort"
	"strings"
)

func GetMeeting(text string) string {
	people := strings.Split(text, ";")
	format := Formating(people)
	var textUpper []string
	for _, name := range format {
		textUpper = append(textUpper, UpperCase(name))
	}
	sort.Strings(textUpper)
	return fmt.Sprintf("%s%s%s%s%s%s%s", textUpper[0], textUpper[1], textUpper[2], textUpper[3], textUpper[4], textUpper[5], textUpper[6])
}

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
