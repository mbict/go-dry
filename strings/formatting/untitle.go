package formatting

import (
	"unicode"
	"unicode/utf8"
)

// UnTitle Make the first charater lower case
func UnTitle(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}
