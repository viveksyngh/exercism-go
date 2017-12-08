package isogram

import "strings"

//IsIsogram check whether a word or phrase is Isogram or not.
func IsIsogram(word string) bool {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charMap := map[string]bool{}

	for _, char := range strings.ToUpper(word) {
		char := string(char)

		if charMap[char] && strings.Contains(alphabet, char) {
			return false
		}
		charMap[char] = true
	}
	return true
}
