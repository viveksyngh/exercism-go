package acronym

import "strings"

// Abbreviate converts a long name to its acronym
func Abbreviate(s string) string {

	alphabets := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result string
	var first string
	first = string(s[0])

	for _, char := range s {
		if strings.Contains(alphabets, string(char)) {
			if first == "" {
				first = string(char)
			}
		} else {
			result = result + strings.ToUpper(string(first))
			first = ""
		}

	}

	if first != "" {
		result = result + strings.ToUpper(string(first))
	}

	return result
}
