package iteration

import "strings"

const repeatCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

func Repeat2(character string, count int) string {
	return strings.Repeat(character, count)
}
