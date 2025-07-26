package util

import "strings"

func TgEscape(txt string) string {
	special := map[rune]bool{
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

	var builder strings.Builder
	builder.Grow(len(txt) * 2)

	for _, r := range txt {
		if special[r] {
			builder.WriteRune('\\')
		}
		builder.WriteRune(r)
	}

	return builder.String()
}
