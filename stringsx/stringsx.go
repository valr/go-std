package stringsx

import (
	"strings"
)

// Return the reversed string.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Split the string into substrings separated by sep and return the slice of substrings.
func SplitAny(s, sep string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return strings.ContainsRune(sep, r)
	})
}
