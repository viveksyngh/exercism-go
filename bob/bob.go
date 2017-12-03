package bob

import (
	"strings"
	"unicode"
)

// IsLetter checks if string contains at least on alphabet.
func IsLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

// Hey gives response to questions asked to bob.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if remark == "" {
		return "Fine. Be that way!"
	} else if remark == strings.ToUpper(remark) && IsLetter(remark) {
		return "Whoa, chill out!"
	} else if strings.HasSuffix(remark, "?") {
		return "Sure."
	}
	return "Whatever."
}
