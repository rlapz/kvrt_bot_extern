package util

import "strings"

var TgSpecial = map[rune]bool{
	'_': true,
	'*': true,
	'[': true,
	']': true,
	'(': true,
	')': true,
	'~': true,
	'`': true,
	'>': true,
	'#': true,
	'+': true,
	'-': true,
	'|': true,
	'{': true,
	'}': true,
	'.': true,
	'!': true,
}

func TgEscape(txt string) string {
	var builder strings.Builder
	builder.Grow(len(txt) * 2)

	for _, r := range txt {
		if TgSpecial[r] {
			builder.WriteRune('\\')
		}

		builder.WriteRune(r)
	}

	return builder.String()
}
